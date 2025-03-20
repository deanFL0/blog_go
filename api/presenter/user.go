package presenter

import (
	"github.com/deanFL0/blog_api_go/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserSuccessResponse(data *entities.User) *fiber.Map {
	user := User{
		ID:       data.ID,
		Name:     data.Name,
		Username: data.Username,
		Password: data.Password,
	}
	return &fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	}
}

func UsersSuccessResponse(data *[]entities.User) *fiber.Map {
	var users []User
	for _, user := range *data {
		users = append(users, User{
			ID:       user.ID,
			Name:     user.Name,
			Username: user.Username,
			Password: user.Password,
		})
	}
	return &fiber.Map{
		"status": true,
		"data":   users,
		"error":  nil,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   "",
		"error":  err.Error(),
	}
}
