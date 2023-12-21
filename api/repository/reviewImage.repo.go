package repository

import (
	"TableTru/infrastructure"
	"TableTru/models"
)

type ReviewImageRepository struct {
	db infrastructure.Database
}

func NewReviewImageRepository(db infrastructure.Database) ReviewImageRepository {
	return ReviewImageRepository{
		db: db,
	}
}

func (c ReviewImageRepository) FindAll(reviewImage models.ReviewImage, keyword string) (*[]models.ReviewImage, int64, error) {
	var reviewImages []models.ReviewImage
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.ReviewImage{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			c.db.DB.Where("reviewImage.review_image_id LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(reviewImage).
		Find(&reviewImages).
		Count(&totalRows).Error
	return &reviewImages, totalRows, err
}

func (c ReviewImageRepository) Find(reviewImage models.ReviewImage) (models.ReviewImage, error) {
	var reviewImages models.ReviewImage
	err := c.db.DB.
		Debug().
		Model(&models.ReviewImage{}).
		Where(&reviewImage).
		Take(&reviewImages).Error
	return reviewImages, err
}

func (c ReviewImageRepository) Create(reviewImage models.ReviewImage) error {
	return c.db.DB.Create(&reviewImage).Error
}

func (c ReviewImageRepository) Update(reviewImage models.ReviewImage) error {
	return c.db.DB.Save(&reviewImage).Error
}

func (c ReviewImageRepository) Delete(reviewImage models.ReviewImage) error {
	return c.db.DB.Delete(&reviewImage).Error
}
