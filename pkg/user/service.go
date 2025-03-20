package user

import (
	"github.com/deanFL0/blog_api_go/pkg/entities"
)

type Service interface {
	InsertUser(user *entities.User) (*entities.User, error)
	FetchUser(ID int) (*entities.User, error)
	FetchUsers() (*[]entities.User, error)
	UpdateUser(ID int, user *entities.User) (*entities.User, error)
	RemoveUser(ID int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertUser(user *entities.User) (*entities.User, error) {
	return s.repository.CreteUser(user)
}

func (s *service) FetchUser(ID int) (*entities.User, error) {
	return s.repository.ReadUser(ID)
}

func (s *service) FetchUsers() (*[]entities.User, error) {
	return s.repository.ReadUsers()
}

func (s *service) UpdateUser(ID int, user *entities.User) (*entities.User, error) {
	return s.repository.UpdateUser(ID, user)
}

func (s *service) RemoveUser(ID int) error {
	return s.repository.DeleteUser(ID)
}
