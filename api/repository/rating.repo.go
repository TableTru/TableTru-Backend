package repository

import (
	"TableTru/infrastructure"
	"TableTru/models"
)

type RatingRepository struct {
	db infrastructure.Database
}

func NewRatingRepository(db infrastructure.Database) RatingRepository {
	return RatingRepository{
		db: db,
	}
}

func (c RatingRepository) FindAll(rating models.Rating, keyword string) (*[]models.Rating, int64, error) {
	var ratings []models.Rating
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.Rating{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			c.db.DB.Where("rating.rating_id LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(rating).
		Find(&ratings).
		Count(&totalRows).Error
	return &ratings, totalRows, err
}

func (c RatingRepository) Find(rating models.Rating) (models.Rating, error) {
	var ratings models.Rating
	err := c.db.DB.
		Debug().
		Model(&models.Rating{}).
		Where(&rating).
		Take(&ratings).Error
	return ratings, err
}

func (c RatingRepository) Create(rating models.Rating) error {
	return c.db.DB.Create(&rating).Error
}

func (c RatingRepository) Update(rating models.Rating) error {
	return c.db.DB.Save(&rating).Error
}

func (c RatingRepository) Delete(rating models.Rating) error {
	return c.db.DB.Delete(&rating).Error
}
