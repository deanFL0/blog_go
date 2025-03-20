package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/deanFL0/blog_api_go/api/presenter"
	"github.com/deanFL0/blog_api_go/pkg/entities"
	"github.com/deanFL0/blog_api_go/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func AddUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		if requestBody.Name == "" || requestBody.Username == "" || requestBody.Password == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(errors.New("please specify title and body")))
		}
		result, err := service.InsertUser(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

func UpdateUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		existingUser, err := service.FetchUser(int(ID))
		if err != nil {
			return c.JSON(presenter.UserErrorResponse(err))
		}

		var requestBody entities.User
		err = c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		if requestBody.Name != "" {
			existingUser.Name = requestBody.Name
		}
		if requestBody.Username != "" {
			existingUser.Username = requestBody.Username
		}

		result, err := service.UpdateUser(ID, existingUser)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

func RemoveUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		err = service.RemoveUser(int(ID))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "Deleted successfully",
			"err":    nil,
		})
	}
}

func GetUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		fetched, err := service.FetchUser(int(ID))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(fetched))
	}
}

func GetUsers(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchUsers()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UsersSuccessResponse(fetched))
	}
}
