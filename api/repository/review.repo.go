package repository

import (
	"TableTru/infrastructure"
	"TableTru/models"
)

type ReviewRepository struct {
	db infrastructure.Database
}

func NewReviewRepository(db infrastructure.Database) ReviewRepository {
	return ReviewRepository{
		db: db,
	}
}

func (c ReviewRepository) FindAll(review models.Review, keyword string) (*[]models.Review, int64, error) {
	var reviews []models.Review
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.Review{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			c.db.DB.Where("review.review_id LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(review).
		Find(&reviews).
		Count(&totalRows).Error
	return &reviews, totalRows, err
}

func (c ReviewRepository) Find(review models.Review) (models.Review, error) {
	var reviews models.Review
	err := c.db.DB.
		Debug().
		Model(&models.Review{}).
		Where(&review).
		Take(&reviews).Error
	return reviews, err
}

func (c ReviewRepository) Create(review models.Review) error {
	return c.db.DB.Create(&review).Error
}

func (c ReviewRepository) Update(review models.Review) error {
	return c.db.DB.Save(&review).Error
}

func (c ReviewRepository) Delete(review models.Review) error {
	return c.db.DB.Delete(&review).Error
}
