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

func (r *UserRepository) SeedData(users []models.User) error {
	result := r.db.DB.Create(users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c UserRepository) FindAll(user models.User, keyword string) (*[]models.User, int64, error) {
	var users []models.User
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.User{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			c.db.DB.Where("user.username LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(user).
		Find(&users).
		Count(&totalRows).Error
	return &users, totalRows, err
}

func (c UserRepository) Find(user models.User) (models.User, error) {
	var users models.User
	err := c.db.DB.
		Debug().
		Model(&models.User{}).
		Where(&user).
		Take(&users).Error
	return users, err
}

func (c UserRepository) Create(user models.User) error {
	return c.db.DB.Create(&user).Error
}

func (c UserRepository) Update(user models.User) error {
	return c.db.DB.Save(&user).Error
}

func (c UserRepository) Delete(user models.User) error {
	return c.db.DB.Delete(&user).Error
}

func (c UserRepository) FindLoginUser(username, password string) (models.User, error) {
	var user models.User
	err := c.db.DB.
		Debug().
		Model(&models.User{}).
		Where("username = ? AND password = ?", username, password).
		Take(&user).Error
	return user, err
}
