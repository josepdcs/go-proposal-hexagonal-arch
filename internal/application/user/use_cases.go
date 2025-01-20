package user

import (
	"context"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/entity"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/port/input/user"
	userrepo "github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/port/output/user"
)

// DefaultFinderAllUseCase use case
type DefaultFinderAllUseCase struct {
	repository userrepo.Repository
}

// NewDefaultFinderAllUseCase creates a new user.DefaultFinderAllUseCase instance
func NewDefaultFinderAllUseCase(repository userrepo.Repository) user.FinderAllUseCase {
	return &DefaultFinderAllUseCase{
		repository: repository,
	}
}

// Find returns all users or an error if something goes wrong
func (u *DefaultFinderAllUseCase) Find(ctx context.Context) ([]entity.User, error) {
	users, err := u.repository.FindAll(ctx)
	return users, err
}

// DefaultFinderByIDUseCase use case
type DefaultFinderByIDUseCase struct {
	repository userrepo.Repository
}

// NewDefaultFinderByIDUseCase creates a new user.DefaultFinderByIDUseCase instance
func NewDefaultFinderByIDUseCase(repository userrepo.Repository) user.FinderByIDUseCase {
	return &DefaultFinderByIDUseCase{
		repository: repository,
	}
}

// Find returns a repository by ID or an error if something goes wrong
func (u *DefaultFinderByIDUseCase) Find(ctx context.Context, id uint) (entity.User, error) {
	return u.repository.FindByID(ctx, id)
}

// DefaultCreatorUseCase defines the use case for creating a repository
type DefaultCreatorUseCase struct {
	repository userrepo.Repository
}

// NewDefaultCreatorUseCase creates a new user.DefaultCreatorUseCase instance
func NewDefaultCreatorUseCase(repository userrepo.Repository) user.CreatorUseCase {
	return &DefaultCreatorUseCase{
		repository: repository,
	}
}

// Create creates a repository and returns the created repository or an error if something goes wrong
func (u *DefaultCreatorUseCase) Create(ctx context.Context, user entity.User) (entity.User, error) {
	return u.repository.Create(ctx, user)
}

// DefaultModifierUseCase defines the use case for modifying a repository
type DefaultModifierUseCase struct {
	repository userrepo.Repository
}

// NewDefaultModifierUseCase creates a new user.DefaultModifierUseCase instance
func NewDefaultModifierUseCase(repository userrepo.Repository) user.ModifierUseCase {
	return &DefaultModifierUseCase{
		repository: repository,
	}
}

// Modify modifies a repository and returns the modified repository or an error if something goes wrong
func (u *DefaultModifierUseCase) Modify(ctx context.Context, user entity.User) (entity.User, error) {
	return u.repository.Modify(ctx, user)
}

// DefaultDeleterUseCase defines the use case for deleting a repository
type DefaultDeleterUseCase struct {
	repository userrepo.Repository
}

// NewDefaultDeleterUseCase creates a new usecase.DefaultDeleterUseCase instance
func NewDefaultDeleterUseCase(repository userrepo.Repository) user.DeleterUseCase {
	return &DefaultDeleterUseCase{
		repository: repository,
	}
}

// Delete deletes a repository and returns an error if something goes wrong
func (u *DefaultDeleterUseCase) Delete(ctx context.Context, user entity.User) error {
	return u.repository.Delete(ctx, user)
}
