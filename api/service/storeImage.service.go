package service

import (
	"TableTru/api/repository"
	"TableTru/models"
)

type StoreImageService struct {
	repository repository.StoreImageRepository
}

func NewStoreImageService(r repository.StoreImageRepository) StoreImageService {
	return StoreImageService{
		repository: r,
	}
}

func (c StoreImageService) FindAllStoreImage(storeImage models.StoreImage, keyword string) (*[]models.StoreImage, int64, error) {
	return c.repository.FindAll(storeImage, keyword)
}

func (c StoreImageService) FindStoreImage(storeImage models.StoreImage) (models.StoreImage, error) {
	return c.repository.Find(storeImage)
}

func (c StoreImageService) CreateStoreImage(storeImage models.StoreImage) error {
	return c.repository.Create(storeImage)
}

func (c StoreImageService) UpdateStoreImage(storeImage models.StoreImage) error {
	return c.repository.Update(storeImage)
}

func (c StoreImageService) DeleteStoreImage(id int64) error {
	var storeImage models.StoreImage
	storeImage.ID = id
	return c.repository.Delete(storeImage)
}
