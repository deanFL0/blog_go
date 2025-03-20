package user

import (
	"github.com/deanFL0/blog_api_go/pkg/entities"
	"gorm.io/gorm"
)

type Repository interface {
	CreteUser(user *entities.User) (*entities.User, error)
	ReadUser(ID int) (*entities.User, error)
	ReadUsers() (*[]entities.User, error)
	UpdateUser(ID int, user *entities.User) (*entities.User, error)
	DeleteUser(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreteUser(user *entities.User) (*entities.User, error) {
	return user, r.db.Create(user).Error
}

func (r *repository) ReadUser(ID int) (*entities.User, error) {
	var user entities.User
	result := r.db.Find(&user, ID)
	return &user, result.Error
}

func (r *repository) ReadUsers() (*[]entities.User, error) {
	var user []entities.User
	result := r.db.Find(&user)
	return &user, result.Error
}

func (r *repository) UpdateUser(ID int, user *entities.User) (*entities.User, error) {
	existingUser, err := r.ReadUser(user.ID)
	if err != nil {
		return nil, err
	}

	existingUser.Name = user.Name
	existingUser.Username = user.Username

	return existingUser, r.db.Save(existingUser).Error
}

func (r *repository) DeleteUser(ID int) error {
	var user entities.User
	err := r.db.Find(&user, ID).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&user).Error
}
