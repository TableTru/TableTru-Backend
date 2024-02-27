package service

import (
	"TableTru/api/repository"
	"TableTru/models"
)

type OpenTimeService struct {
	repository repository.OpenTimeRepository
}

func NewOpenTimeService(r repository.OpenTimeRepository) OpenTimeService {
	return OpenTimeService{
		repository: r,
	}
}

func (c OpenTimeService) FindAllOpenTime(openTime models.OpenTime, keyword string) (*[]models.OpenTime, int64, error) {
	return c.repository.FindAll(openTime, keyword)
}

func (c OpenTimeService) FindOpenTime(openTime models.OpenTime) (models.OpenTime, error) {
	return c.repository.Find(openTime)
}

func (c OpenTimeService) CreateOpenTime(openTime models.OpenTime) error {
	return c.repository.Create(openTime)
}

func (c OpenTimeService) UpdateOpenTime(openTime models.OpenTime) error {
	return c.repository.Update(openTime)
}

func (c OpenTimeService) DeleteOpenTime(id int64) error {
	var openTime models.OpenTime
	openTime.ID = id
	return c.repository.Delete(openTime)
}
