package service

import (
	"TableTru/api/repository"
	"TableTru/models"
)

type CategoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return CategoryService{
		repository: r,
	}
}

func (c CategoryService) FindAllCategory(category models.Category, keyword string) (*[]models.Category, int64, error) {
	return c.repository.FindAll(category, keyword)
}

func (c CategoryService) FindCategory(category models.Category) (models.Category, error) {
	return c.repository.Find(category)
}

func (c CategoryService) CreateCategory(category models.Category) error {
	return c.repository.Create(category)
}

func (c CategoryService) UpdateCategory(category models.Category) error {
	return c.repository.Update(category)
}

func (c CategoryService) DeleteCategory(id int64) error {
	var category models.Category
	category.ID = id
	return c.repository.Delete(category)
}
