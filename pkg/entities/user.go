package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
