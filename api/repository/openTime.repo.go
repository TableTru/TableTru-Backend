package repository

import (
	"TableTru/infrastructure"
	"TableTru/models"
)

type OpenTimeRepository struct {
	db infrastructure.Database
}

func NewOpenTimeRepository(db infrastructure.Database) OpenTimeRepository {
	return OpenTimeRepository{
		db: db,
	}
}

func (r *OpenTimeRepository) SeedData(openTimes []models.OpenTime) error {
	result := r.db.DB.Create(openTimes)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c OpenTimeRepository) FindAll(openTime models.OpenTime, keyword string) (*[]models.OpenTime, int64, error) {
	var openTimes []models.OpenTime
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.OpenTime{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			c.db.DB.Where("openTime.openTime_id LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(openTime).
		Find(&openTimes).
		Count(&totalRows).Error
	return &openTimes, totalRows, err
}

func (c OpenTimeRepository) Find(openTime models.OpenTime) (models.OpenTime, error) {
	var openTimes models.OpenTime
	err := c.db.DB.
		Debug().
		Model(&models.OpenTime{}).
		Where(&openTime).
		Take(&openTimes).Error
	return openTimes, err
}

func (c OpenTimeRepository) Create(openTime models.OpenTime) error {
	return c.db.DB.Create(&openTime).Error
}

func (c OpenTimeRepository) Update(openTime models.OpenTime) error {
	return c.db.DB.Save(&openTime).Error
}

func (c OpenTimeRepository) Delete(openTime models.OpenTime) error {
	return c.db.DB.Delete(&openTime).Error
}
