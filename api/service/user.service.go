package service

import (
	"TableTru/api/repository"
	"TableTru/models"
)

// PostService PostService struct
type UserService struct {
	repository repository.UserRepository
}

// NewPostService : returns the PostService struct instance
func NewUserService(r repository.UserRepository) UserService {
	return UserService{
		repository: r,
	}
}

func (u UserService) CreateUser(user models.User) error {
	return u.repository.CreateUser(user)
}

func (u UserService) FindAllUser(user models.User, keyword string) (*[]models.User, int64, error) {
	return u.repository.FindAllUser(user, keyword)
}
