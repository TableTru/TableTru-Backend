package service

import (
	"TableTru/api/repository"
	"TableTru/models"
	"fmt"
	"sort"
	"time"
)

type TableBookingService struct {
	repository repository.TableBookingRepository
}


type ByBookingObject []models.TimeObject

func (a ByBookingObject) Len() int           { return len(a) }
func (a ByBookingObject) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByBookingObject) Less(i, j int) bool { return a[i].StartTime.Before(a[j].StartTime) }

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

func (c TableBookingService) FindAllUserBookingByStatus(tableBooking models.TableBooking) (*[]models.TableBooking, int64, error) {
	return c.repository.FindAllUserBookingByStatus(tableBooking)
}

func (c TableBookingService) CheckBooking(tableBooking models.TableBooking, keyword string, maxCount int) (*[]models.TimeObject, int64, error) {
	tableBookings, totalRows, err := c.repository.FindAll(tableBooking, keyword)
	if err != nil {
		return nil, 0, err
	}
	
	var timeObjects []models.TimeObject
	for _, booking := range *tableBookings {
		timeObjects = append(timeObjects, models.TimeObject{StartTime: booking.BookingTime, EndTime: booking.BookingTime.Add(time.Hour)})
	}

	var disableTimeObjects []models.TimeObject

	sort.Sort(ByBookingObject(timeObjects))
	for i := 0; i < len(timeObjects); i++ {
		startTime := timeObjects[i].StartTime
		endTime := timeObjects[i].EndTime
		count := 1

		for j := i + 1; j < len(timeObjects); j++ {
			if (timeObjects[j].StartTime.After(startTime) || timeObjects[j].StartTime.Equal(startTime)) && timeObjects[j].StartTime.Before(endTime) {
				count++
			}
		}
		
		if count >= maxCount {
            found := false
        for _, obj := range disableTimeObjects {
            if obj.StartTime.Equal(startTime) {
                found = true
                break
            }
        }
        if !found {
            fmt.Printf("Number of overlapping time ranges for %s - %s: %d\n", startTime, endTime, count)
            disableTimeObjects = append(disableTimeObjects, models.TimeObject{StartTime: startTime, EndTime: endTime})
        }
        }
	}

	return &disableTimeObjects, totalRows, nil
}
