package repository

import (
	"TableTru/infrastructure"
	"TableTru/models"
)

type UserRepository struct {
	db infrastructure.Database
}

func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (u UserRepository) CreateUser(user models.User) error {
	return u.db.DB.Create(&user).Error
}

func (u UserRepository) FindAllUser(user models.User, keyword string) (*[]models.User, int64, error) {
	var users []models.User
	var totalRows int64 = 0

	queryBuider := u.db.DB.Order("created_at desc").Model(&models.User{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			u.db.DB.Where("user.name LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(user).
		Find(&users).
		Count(&totalRows).Error
	return &users, totalRows, err
}
