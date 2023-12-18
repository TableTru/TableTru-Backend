package repository

import (
	"TableTru/infrastructure"
	"TableTru/models"
)

type CategoryRepository struct {
	db infrastructure.Database
}

func NewCategoryRepository(db infrastructure.Database) CategoryRepository {
	return CategoryRepository{
		db: db,
	}
}

func (c CategoryRepository) FindAll(category models.Category, keyword string) (*[]models.Category, int64, error) {
	var categories []models.Category
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.Category{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			c.db.DB.Where("category.category_name LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(category).
		Find(&categories).
		Count(&totalRows).Error
	return &categories, totalRows, err
}

func (c CategoryRepository) Find(category models.Category) (models.Category, error) {
	var categories models.Category
	err := c.db.DB.
		Debug().
		Model(&models.Category{}).
		Where(&category).
		Take(&categories).Error
	return categories, err
}

func (c CategoryRepository) Create(category models.Category) error {
	return c.db.DB.Create(&category).Error
}

func (c CategoryRepository) Update(category models.Category) error {
	return c.db.DB.Save(&category).Error
}

func (c CategoryRepository) Delete(category models.Category) error {
	return c.db.DB.Delete(&category).Error
}
