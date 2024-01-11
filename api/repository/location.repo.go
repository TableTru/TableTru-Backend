package repository

import (
	"TableTru/infrastructure"
	"TableTru/models"
)

type LocationRepository struct {
	db infrastructure.Database
}

func NewLocationRepository(db infrastructure.Database) LocationRepository {
	return LocationRepository{
		db: db,
	}
}

func (r *LocationRepository) SeedData(locations []models.Location) error {
	result := r.db.DB.Create(locations)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c LocationRepository) FindAll(location models.Location, keyword string) (*[]models.Location, int64, error) {
	var locations []models.Location
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.Location{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			c.db.DB.Where("location.location_name LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(location).
		Find(&locations).
		Count(&totalRows).Error
	return &locations, totalRows, err
}

func (c LocationRepository) Find(location models.Location) (models.Location, error) {
	var locations models.Location
	err := c.db.DB.
		Debug().
		Model(&models.Location{}).
		Where(&location).
		Take(&locations).Error
	return locations, err
}

func (c LocationRepository) Create(location models.Location) error {
	return c.db.DB.Create(&location).Error
}

func (c LocationRepository) Update(location models.Location) error {
	return c.db.DB.Save(&location).Error
}

func (c LocationRepository) Delete(location models.Location) error {
	return c.db.DB.Delete(&location).Error
}
