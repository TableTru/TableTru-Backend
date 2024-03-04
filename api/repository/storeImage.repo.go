package repository

import (
	"TableTru/infrastructure"
	"TableTru/models"
)

type StoreImageRepository struct {
	db infrastructure.Database
}

func NewStoreImageRepository(db infrastructure.Database) StoreImageRepository {
	return StoreImageRepository{
		db: db,
	}
}

func (c StoreImageRepository) FindAll(storeImage models.StoreImage, keyword string) (*[]models.StoreImage, int64, error) {
	var storeImages []models.StoreImage
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.StoreImage{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			c.db.DB.Where("storeImage.store_image_id LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(storeImage).
		Find(&storeImages).
		Count(&totalRows).Error
	return &storeImages, totalRows, err
}

func (c StoreImageRepository) Find(storeImage models.StoreImage) (models.StoreImage, error) {
	var storeImages models.StoreImage
	err := c.db.DB.
		Debug().
		Model(&models.StoreImage{}).
		Where(&storeImage).
		Take(&storeImages).Error
	return storeImages, err
}

func (c StoreImageRepository) Create(storeImage models.StoreImage) error {
	return c.db.DB.Create(&storeImage).Error
}

func (c StoreImageRepository) Update(storeImage models.StoreImage) error {
	return c.db.DB.Save(&storeImage).Error
}

func (c StoreImageRepository) Delete(storeImage models.StoreImage) error {
	return c.db.DB.Delete(&storeImage).Error
}

func (c StoreImageRepository) FindStoreImageByType(storeImage models.StoreImage) (*[]models.StoreImage, int64, error) {
	var storeImages []models.StoreImage
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.StoreImage{})

	// Search parameter
	// if storeType != "" {
	// 	queryKeyword := "%" + storeType + "%"
	// 	queryBuider = queryBuider.Where(
	// 		c.db.DB.Where("storeImage.store_image_type LIKE ? AND storeImage.store_id LIKE ? ", queryKeyword, id))
	// }

	err := queryBuider.
		Where(storeImage).
		Find(&storeImages).
		Count(&totalRows).Error
	return &storeImages, totalRows, err
}