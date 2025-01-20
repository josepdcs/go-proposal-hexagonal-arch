package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/entity"
	domerrors "github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/model/errors"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/port/input/user"
	"github.com/pkg/errors"
)

const errBuildResponseTxt = "cannot build response"

// UserAPI encapsulates the user use cases.
type UserAPI struct {
	finderAllUseCase  user.FinderAllUseCase
	finderByIDUseCase user.FinderByIDUseCase
	creatorUseCase    user.CreatorUseCase
	modifierUseCase   user.ModifierUseCase
	deleterUseCase    user.DeleterUseCase
}

type UserDTO struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

// toEntityUser converts a UserDTO to an entity.User
func (u UserDTO) toEntityUser() entity.User {
	return entity.User{
		ID:      u.ID,
		Name:    u.Name,
		Surname: u.Surname,
	}
}

// toUserDTO concerts entity.User to UserDTO
func toUserDTO(u entity.User) UserDTO {
	return UserDTO{
		ID:      u.ID,
		Name:    u.Name,
		Surname: u.Surname,
	}
}

// NewUserAPI creates a new UserAPI.
func NewUserAPI(
	finderAllUseCase user.FinderAllUseCase,
	finderByIDUseCase user.FinderByIDUseCase,
	creatorUseCase user.CreatorUseCase,
	modifierUseCase user.ModifierUseCase,
	deleterUseCase user.DeleterUseCase,
) *UserAPI {
	return &UserAPI{
		finderAllUseCase:  finderAllUseCase,
		finderByIDUseCase: finderByIDUseCase,
		creatorUseCase:    creatorUseCase,
		modifierUseCase:   modifierUseCase,
		deleterUseCase:    deleterUseCase,
	}
}

// FindAll godoc
// @summary Get all users
// @description Get all users
// @tags users
// @security ApiKeyAuth
// @id FindAll
// @produce json
// @Router /api/users [get]
// @response 200 {object} []UserDTO "OK"
func (h *UserAPI) FindAll(c *fiber.Ctx) error {
	users, err := h.finderAllUseCase.Find(c.UserContext())

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		response := make([]UserDTO, 0, len(users))
		for _, u := range users {
			response = append(response, toUserDTO(u))
		}
		return c.JSON(response)
	}
}

// FindByID godoc
// @summary Get a user by ID
// @description Get a user by ID
// @tags users
// @security ApiKeyAuth
// @id FindByID
// @produce json
// @param id path int true "User ID"
// @Router /api/users/{id} [get]
// @response 200 {object} UserDTO "OK"
func (h *UserAPI) FindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, "cannot parse id"))
	}

	u, err := h.finderByIDUseCase.Find(c.UserContext(), uint(id))

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		return c.JSON(toUserDTO(u))
	}
}

// Create godoc
// @summary Create a user
// @description Create a user
// @tags users
// @security ApiKeyAuth
// @id Create
// @accept json
// @produce json
// @param user body entity.User true "entity.User"
// @Router /api/users [post]
// @response 200 {object} UserDTO "OK"
func (h *UserAPI) Create(c *fiber.Ctx) error {
	var userDTO UserDTO

	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	u, err := h.creatorUseCase.Create(c.UserContext(), userDTO.toEntityUser())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, "Cannot create user: "+err.Error()))
	} else {
		return c.Status(fiber.StatusCreated).JSON(toUserDTO(u))
	}
}

// Modify godoc
// @summary Modify a user
// @description Modify a user
// @tags users
// @security ApiKeyAuth
// @id Modify
// @accept json
// @produce json
// @param user body entity.User true "entity.User"
// @Router /api/users [put]
// @response 200 {object} UserDTO "OK"
func (h *UserAPI) Modify(c *fiber.Ctx) error {
	var userDTO UserDTO

	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	u, err := h.modifierUseCase.Modify(c.UserContext(), userDTO.toEntityUser())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, "Cannot modify user: "+err.Error()))
	} else {
		return c.JSON(toUserDTO(u))
	}
}

// Delete godoc
// @summary Delete a user
// @description Delete a user
// @tags users
// @security ApiKeyAuth
// @id Delete
// @param id path int true "User ID"
// @Router /api/users/{id} [delete]
// @response 200 {object} UserDTO "OK"
func (h *UserAPI) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, "cannot parse id"))
	}

	u, err := h.finderByIDUseCase.Find(c.UserContext(), uint(id))

	if err != nil {
		if errors.Is(err, domerrors.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.NewError(fiber.StatusNotFound, "User not found"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	if u == (entity.User{}) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.NewError(fiber.StatusNotFound, "User not found"))
	}

	err = h.deleterUseCase.Delete(c.UserContext(), u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, "Cannot delete user: "+err.Error()))
	}

	return c.SendStatus(fiber.StatusNoContent)
}
