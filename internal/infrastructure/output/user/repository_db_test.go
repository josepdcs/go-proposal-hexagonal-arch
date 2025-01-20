package user

import (
	"context"
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/entity"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/port/output/user"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/output/testutil"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestDBRepository_FindAll(t *testing.T) {
	tests := []struct {
		name  string
		given func() (user.Repository, sqlmock.Sqlmock)
		when  func(r user.Repository) ([]entity.User, error)
		then  func(sqlmock.Sqlmock, []entity.User, error)
	}{
		{
			name: "should find all users",
			given: func() (user.Repository, sqlmock.Sqlmock) {
				db, mock, err := testutil.NewMockPostgresSqlDB()
				if err != nil {
					t.Fatal(err)
				}

				rows := sqlmock.NewRows([]string{"id", "name", "surname"}).
					AddRow(1, "John", "Doe").
					AddRow(2, "Jane", "Doe").
					AddRow(3, "Alice", "Smith")

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL`)).
					WillReturnRows(rows)

				return NewDBRepository(db), mock
			},
			when: func(r user.Repository) ([]entity.User, error) {
				return r.FindAll(context.Background())
			},
			then: func(mock sqlmock.Sqlmock, users []entity.User, err error) {
				assert.NoError(t, err)
				assert.Len(t, users, 3)
				assert.Equal(t, "John", users[0].Name)
				assert.Equal(t, "Doe", users[0].Surname)
				assert.Equal(t, "Jane", users[1].Name)
				assert.Equal(t, "Doe", users[1].Surname)
				assert.Equal(t, "Alice", users[2].Name)
				assert.Equal(t, "Smith", users[2].Surname)

				assert.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "should not find users",
			given: func() (user.Repository, sqlmock.Sqlmock) {
				db, mock, err := testutil.NewMockPostgresSqlDB()
				if err != nil {
					t.Fatal(err)
				}

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL`)).
					WillReturnError(errors.New("not found"))

				return NewDBRepository(db), mock
			},
			when: func(r user.Repository) ([]entity.User, error) {
				return r.FindAll(context.Background())
			},
			then: func(mock sqlmock.Sqlmock, users []entity.User, err error) {
				assert.Error(t, err)
				assert.Len(t, users, 0)

				assert.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			repo, mock := tt.given()

			// When
			users, err := tt.when(repo)

			// Then
			tt.then(mock, users, err)
		})
	}
}

func TestDBRepository_FindByID(t *testing.T) {
	tests := []struct {
		name  string
		given func() (user.Repository, sqlmock.Sqlmock)
		when  func(r user.Repository) (entity.User, error)
		then  func(sqlmock.Sqlmock, entity.User, error)
	}{
		{
			name: "should find user by ID",
			given: func() (user.Repository, sqlmock.Sqlmock) {
				db, mock, err := testutil.NewMockPostgresSqlDB()
				if err != nil {
					t.Fatal(err)
				}

				rows := sqlmock.NewRows([]string{"id", "name", "surname"}).
					AddRow(1, "John", "Doe").
					AddRow(2, "Jane", "Doe").
					AddRow(3, "Alice", "Smith")

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
					WithArgs(1, 1).
					WillReturnRows(rows)

				return NewDBRepository(db), mock
			},
			when: func(r user.Repository) (entity.User, error) {
				return r.FindByID(context.Background(), 1)
			},
			then: func(mock sqlmock.Sqlmock, user entity.User, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "John", user.Name)
				assert.Equal(t, "Doe", user.Surname)

				assert.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "should not find user by ID",
			given: func() (user.Repository, sqlmock.Sqlmock) {
				db, mock, err := testutil.NewMockPostgresSqlDB()
				if err != nil {
					t.Fatal(err)
				}

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
					WithArgs(1, 1).
					WillReturnError(errors.New("not found"))

				return NewDBRepository(db), mock
			},
			when: func(r user.Repository) (entity.User, error) {
				return r.FindByID(context.Background(), 1)
			},
			then: func(mock sqlmock.Sqlmock, user entity.User, err error) {
				assert.Error(t, err)
				assert.Empty(t, user)

				assert.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			repo, mock := tt.given()

			// When
			users, err := tt.when(repo)

			// Then
			tt.then(mock, users, err)
		})
	}
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestDBRepository_Create(t *testing.T) {
	tests := []struct {
		name  string
		given func() (user.Repository, sqlmock.Sqlmock)
		when  func(r user.Repository) (entity.User, error)
		then  func(sqlmock.Sqlmock, entity.User, error)
	}{
		{
			name: "should create user",
			given: func() (user.Repository, sqlmock.Sqlmock) {
				// here we create a new mock database for MySQL due to the limitations of go-sqlmock with PostgresSQL
				// see https://github.com/DATA-DOG/go-sqlmock/issues/118
				db, mock, err := testutil.NewMockMySqlDB()
				if err != nil {
					t.Fatal(err)
				}

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`surname`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?)")).
					WithArgs("John", "Doe", AnyTime{}, AnyTime{}, nil).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				/*mock.ExpectQuery(`INSERT INTO "users" ("name","surname") VALUES ($1,$2)  RETURNING`).
				WithArgs("John", "Doe").
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))*/

				return NewDBRepository(db), mock
			},
			when: func(r user.Repository) (entity.User, error) {
				return r.Create(context.Background(), entity.User{Name: "John", Surname: "Doe"})
			},
			then: func(mock sqlmock.Sqlmock, user entity.User, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "John", user.Name)
				assert.Equal(t, "Doe", user.Surname)

				assert.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "should not create user",
			given: func() (user.Repository, sqlmock.Sqlmock) {
				// here we create a new mock database for MySQL due to the limitations of go-sqlmock with PostgresSQL
				// see https://github.com/DATA-DOG/go-sqlmock/issues/118
				db, mock, err := testutil.NewMockMySqlDB()
				if err != nil {
					t.Fatal(err)
				}

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`surname`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?)")).
					WithArgs("John", "Doe", AnyTime{}, AnyTime{}, nil).
					WillReturnError(errors.New("failed to create user"))
				mock.ExpectRollback()

				/*mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("name","surname") VALUES ($1,$2) RETURNING "id"`)).
					WithArgs("John", "Doe").
					WillReturnError(errors.New("failed to create user"))
				mock.ExpectRollback()*/

				return NewDBRepository(db), mock
			},
			when: func(r user.Repository) (entity.User, error) {
				return r.Create(context.Background(), entity.User{Name: "John", Surname: "Doe"})
			},
			then: func(mock sqlmock.Sqlmock, user entity.User, err error) {
				assert.Error(t, err)
				assert.Empty(t, user)

				assert.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			repo, mock := tt.given()

			// When
			users, err := tt.when(repo)

			// Then
			tt.then(mock, users, err)
		})
	}
}

func TestDBRepository_Modify(t *testing.T) {
	tests := []struct {
		name  string
		given func() (user.Repository, sqlmock.Sqlmock)
		when  func(r user.Repository) (entity.User, error)
		then  func(sqlmock.Sqlmock, entity.User, error)
	}{
		{
			name: "should modify user",
			given: func() (user.Repository, sqlmock.Sqlmock) {
				// here we create a new mock database for MySQL due to the limitations of go-sqlmock with PostgresSQL
				// see https://github.com/DATA-DOG/go-sqlmock/issues/118
				db, mock, err := testutil.NewMockMySqlDB()
				if err != nil {
					t.Fatal(err)
				}

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `name`=?,`surname`=?,`created_at`=?,`updated_at`=?,`deleted_at`=? WHERE `users`.`deleted_at` IS NULL AND `id` = ?")).
					WithArgs("John", "Doe", AnyTime{}, AnyTime{}, nil, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				return NewDBRepository(db), mock
			},
			when: func(r user.Repository) (entity.User, error) {
				return r.Modify(context.Background(), entity.User{ID: 1, Name: "John", Surname: "Doe"})
			},
			then: func(mock sqlmock.Sqlmock, user entity.User, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "John", user.Name)
				assert.Equal(t, "Doe", user.Surname)

				assert.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "should not modify user",
			given: func() (user.Repository, sqlmock.Sqlmock) {
				// here we create a new mock database for MySQL due to the limitations of go-sqlmock with PostgresSQL
				// see https://github.com/DATA-DOG/go-sqlmock/issues/118
				db, mock, err := testutil.NewMockMySqlDB()
				if err != nil {
					t.Fatal(err)
				}

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `name`=?,`surname`=?,`created_at`=?,`updated_at`=?,`deleted_at`=? WHERE `users`.`deleted_at` IS NULL AND `id` = ?")).
					WithArgs("John", "Doe", AnyTime{}, AnyTime{}, nil, 1).
					WillReturnError(errors.New("failed to create user"))
				mock.ExpectRollback()

				return NewDBRepository(db), mock
			},
			when: func(r user.Repository) (entity.User, error) {
				return r.Modify(context.Background(), entity.User{ID: 1, Name: "John", Surname: "Doe"})
			},
			then: func(mock sqlmock.Sqlmock, user entity.User, err error) {
				assert.Error(t, err)
				assert.Empty(t, user)

				assert.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			repo, mock := tt.given()

			// When
			users, err := tt.when(repo)

			// Then
			tt.then(mock, users, err)
		})
	}
}

func TestDBRepository_Delete(t *testing.T) {
	tests := []struct {
		name  string
		given func() (user.Repository, sqlmock.Sqlmock)
		when  func(r user.Repository) error
		then  func(sqlmock.Sqlmock, error)
	}{
		{
			name: "should find user by ID",
			given: func() (user.Repository, sqlmock.Sqlmock) {
				db, mock, err := testutil.NewMockPostgresSqlDB()
				if err != nil {
					t.Fatal(err)
				}

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "deleted_at"=$1 WHERE "users"."id" = $2 AND "users"."deleted_at" IS NULL`)).
					WithArgs(AnyTime{}, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				return NewDBRepository(db), mock
			},
			when: func(r user.Repository) error {
				return r.Delete(context.Background(), entity.User{ID: 1, Name: "John", Surname: "Doe"})
			},
			then: func(mock sqlmock.Sqlmock, err error) {
				assert.NoError(t, err)

				assert.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "should not find user by ID",
			given: func() (user.Repository, sqlmock.Sqlmock) {
				db, mock, err := testutil.NewMockPostgresSqlDB()
				if err != nil {
					t.Fatal(err)
				}

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "deleted_at"=$1 WHERE "users"."id" = $2 AND "users"."deleted_at" IS NULL`)).
					WithArgs(AnyTime{}, 1).
					WillReturnError(errors.New("not found"))
				mock.ExpectRollback()

				return NewDBRepository(db), mock
			},
			when: func(r user.Repository) error {
				return r.Delete(context.Background(), entity.User{ID: 1, Name: "John", Surname: "Doe"})
			},
			then: func(mock sqlmock.Sqlmock, err error) {
				assert.Error(t, err)

				assert.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			repo, mock := tt.given()

			// When
			err := tt.when(repo)

			// Then
			tt.then(mock, err)
		})
	}
}
