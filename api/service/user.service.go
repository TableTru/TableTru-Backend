package service

import (
	"TableTru/api/repository"
	"TableTru/models"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return UserService{
		repository: r,
	}
}

func (c UserService) FindAllUser(user models.User, keyword string) (*[]models.User, int64, error) {
	return c.repository.FindAll(user, keyword)
}

func (c UserService) FindUser(user models.User) (models.User, error) {
	return c.repository.Find(user)
}

func (c UserService) CreateUser(user models.User) error {
	return c.repository.Create(user)
}

func (c UserService) UpdateUser(user models.User) error {
	return c.repository.Update(user)
}

func (c UserService) DeleteUser(id int64) error {
	var user models.User
	user.ID = id
	return c.repository.Delete(user)
}

func (c UserService) FindLoginUser(username, password string) (models.User, error) {
	return c.repository.FindLoginUser(username, password)
}
