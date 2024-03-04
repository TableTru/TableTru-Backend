package repository

import (
	"TableTru/infrastructure"
	"TableTru/models"
)

type TableBookingRepository struct {
	db infrastructure.Database
}

func NewTableBookingRepository(db infrastructure.Database) TableBookingRepository {
	return TableBookingRepository{
		db: db,
	}
}

func (c TableBookingRepository) FindAll(tableBooking models.TableBooking, keyword string) (*[]models.TableBooking, int64, error) {
	var tableBookings []models.TableBooking
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.TableBooking{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			c.db.DB.Where("tableBooking.table_booking_id LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(tableBooking).
		Find(&tableBookings).
		Count(&totalRows).Error
	return &tableBookings, totalRows, err
}

func (c TableBookingRepository) Find(tableBooking models.TableBooking) (models.TableBooking, error) {
	var tableBookings models.TableBooking
	err := c.db.DB.
		Debug().
		Model(&models.TableBooking{}).
		Where(&tableBooking).
		Take(&tableBookings).Error
	return tableBookings, err
}

func (c TableBookingRepository) Create(tableBooking models.TableBooking) error {
	return c.db.DB.Create(&tableBooking).Error
}

func (c TableBookingRepository) Update(tableBooking models.TableBooking) error {
	return c.db.DB.Save(&tableBooking).Error
}

func (c TableBookingRepository) Delete(tableBooking models.TableBooking) error {
	return c.db.DB.Delete(&tableBooking).Error
}

func (c TableBookingRepository) FindAllUserBookingByStatus(tableBooking models.TableBooking) (*[]models.TableBooking, int64, error) {
	var tableBookings []models.TableBooking
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.TableBooking{})

	// Search parameter
	// if status != "" {
	// 	queryKeyword := "%" + status + "%"
	// 	queryBuider = queryBuider.Where(
	// 		c.db.DB.Where("tableBooking.table_booking_status LIKE ? AND tableBooking.table_booking_id LIKE ? ", queryKeyword, id))
	// }

	err := queryBuider.
		Where(tableBooking).
		Find(&tableBookings).
		Count(&totalRows).Error
	return &tableBookings, totalRows, err
}