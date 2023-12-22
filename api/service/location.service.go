package service

import (
	"TableTru/api/repository"
	"TableTru/models"
)

type LocationService struct {
	repository repository.LocationRepository
}

func NewLocationService(r repository.LocationRepository) LocationService {
	return LocationService{
		repository: r,
	}
}

func (c LocationService) FindAllLocation(location models.Location, keyword string) (*[]models.Location, int64, error) {
	return c.repository.FindAll(location, keyword)
}

func (c LocationService) FindLocation(location models.Location) (models.Location, error) {
	return c.repository.Find(location)
}

func (c LocationService) CreateLocation(location models.Location) error {
	return c.repository.Create(location)
}

func (c LocationService) UpdateLocation(location models.Location) error {
	return c.repository.Update(location)
}

func (c LocationService) DeleteLocation(id int64) error {
	var location models.Location
	location.ID = id
	return c.repository.Delete(location)
}
