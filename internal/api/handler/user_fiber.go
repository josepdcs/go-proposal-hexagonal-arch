package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/entity"
	domerrors "github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/errors"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/domain/usecase"
	"github.com/pkg/errors"
)

const errBuildResponseTxt = "cannot build response"

// UserAPI encapsulates the user use cases.
type UserAPI struct {
	finderAll  usecase.UserFinderAll
	finderByID usecase.UserFinderByID
	creator    usecase.UserCreator
	modifier   usecase.UserModifier
	deleter    usecase.UserDeleter
}

type Response struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// NewUserAPI creates a new UserAPI.
func NewUserAPI(
	finderAll usecase.UserFinderAll,
	finderByID usecase.UserFinderByID,
	creator usecase.UserCreator,
	modifier usecase.UserModifier,
	deleter usecase.UserDeleter,
) *UserAPI {
	return &UserAPI{
		finderAll:  finderAll,
		finderByID: finderByID,
		creator:    creator,
		modifier:   modifier,
		deleter:    deleter,
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
// @response 200 {object} []Response "OK"
func (h *UserAPI) FindAll(c *fiber.Ctx) error {
	users, err := h.finderAll.Find(c.UserContext())

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		response := make([]Response, 0, len(users))
		err = copier.Copy(&response, &users)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
				Error: errBuildResponseTxt,
			})
		}
		return c.JSON(response)
	}
}

func (h *UserAPI) FindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "cannot parse id",
		})
	}

	user, err := h.finderByID.Find(c.UserContext(), uint(id))

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		response := Response{}
		err = copier.Copy(&response, &user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
				Error: errBuildResponseTxt,
			})
		}

		return c.JSON(response)
	}
}

func (h *UserAPI) Create(c *fiber.Ctx) error {
	var user entity.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	user, err := h.creator.Create(c.UserContext(), user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Cannot create user: " + err.Error(),
		})
	} else {
		response := Response{}
		err = copier.Copy(&response, &user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
				Error: errBuildResponseTxt,
			})
		}

		return c.JSON(response)
	}
}

func (h *UserAPI) Modify(c *fiber.Ctx) error {
	var user entity.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	user, err := h.modifier.Modify(c.UserContext(), user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Cannot modify user: " + err.Error(),
		})
	} else {
		response := Response{}
		err = copier.Copy(&response, &user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
				Error: errBuildResponseTxt,
			})
		}

		return c.JSON(response)
	}
}

func (h *UserAPI) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "cannot parse id",
		})
	}

	user, err := h.finderByID.Find(c.UserContext(), uint(id))

	if err != nil {
		if errors.Is(err, domerrors.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
				Error: "User is not registered yet",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	if user == (entity.User{}) {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: "User is not registered yet",
		})
	}

	err = h.deleter.Delete(c.UserContext(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Cannot delete user: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
