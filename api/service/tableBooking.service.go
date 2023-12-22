package service

import (
	"TableTru/api/repository"
	"TableTru/models"
)

type TableBookingService struct {
	repository repository.TableBookingRepository
}

func NewTableBookingService(r repository.TableBookingRepository) TableBookingService {
	return TableBookingService{
		repository: r,
	}
}

func (c TableBookingService) FindAllTableBooking(tableBooking models.TableBooking, keyword string) (*[]models.TableBooking, int64, error) {
	return c.repository.FindAll(tableBooking, keyword)
}

func (c TableBookingService) FindTableBooking(tableBooking models.TableBooking) (models.TableBooking, error) {
	return c.repository.Find(tableBooking)
}

func (c TableBookingService) CreateTableBooking(tableBooking models.TableBooking) error {
	return c.repository.Create(tableBooking)
}

func (c TableBookingService) UpdateTableBooking(tableBooking models.TableBooking) error {
	return c.repository.Update(tableBooking)
}

func (c TableBookingService) DeleteTableBooking(id int64) error {
	var tableBooking models.TableBooking
	tableBooking.ID = id
	return c.repository.Delete(tableBooking)
}
