package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/entity"
	"github.com/thnkrn/go-gin-clean-arch/internal/domain/usecase"
)

const errBuildResponseTxt = "cannot build response"

// UserAPI encapsulates the user use cases.
type UserAPI struct {
	findAll  usecase.UserFinderAll
	findByID usecase.UserFinderByID
	create   usecase.UserCreator
	modify   usecase.UserModifier
	delete   usecase.UserDeleter
}

type Response struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

// NewUserAPI creates a new UserAPI.
func NewUserAPI(
	findAll usecase.UserFinderAll,
	findByID usecase.UserFinderByID,
	create usecase.UserCreator,
	modify usecase.UserModifier,
	delete usecase.UserDeleter,
) *UserAPI {
	return &UserAPI{
		findAll:  findAll,
		findByID: findByID,
		create:   create,
		modify:   modify,
		delete:   delete,
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
func (h *UserAPI) FindAll(c *gin.Context) {
	users, err := h.findAll.Find(c.Request.Context())

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := make([]Response, 0, len(users))
		err = copier.Copy(&response, &users)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errBuildResponseTxt,
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func (h *UserAPI) FindByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	user, err := h.findByID.Find(c.Request.Context(), uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		err = copier.Copy(&response, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errBuildResponseTxt,
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func (h *UserAPI) Create(c *gin.Context) {
	var user entity.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := h.create.Create(c.Request.Context(), user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		err = copier.Copy(&response, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errBuildResponseTxt,
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func (h *UserAPI) Modify(c *gin.Context) {
	var user entity.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := h.modify.Modify(c.Request.Context(), user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		err = copier.Copy(&response, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errBuildResponseTxt,
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func (h *UserAPI) Delete(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot parse id",
		})
		return
	}

	ctx := c.Request.Context()
	user, err := h.findByID.Find(ctx, uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if user == (entity.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User is not booking yet",
		})
		return
	}

	err = h.delete.Delete(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot delete user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "UserAPI is deleted successfully"})
}
