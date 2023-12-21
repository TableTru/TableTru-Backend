package repository

import (
	"TableTru/infrastructure"
	"TableTru/models"
)

type PromotionRepository struct {
	db infrastructure.Database
}

func NewPromotionRepository(db infrastructure.Database) PromotionRepository {
	return PromotionRepository{
		db: db,
	}
}

func (c PromotionRepository) FindAll(promotion models.Promotion, keyword string) (*[]models.Promotion, int64, error) {
	var promotions []models.Promotion
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.Promotion{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			c.db.DB.Where("promotion.promotion_name LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(promotion).
		Find(&promotions).
		Count(&totalRows).Error
	return &promotions, totalRows, err
}

func (c PromotionRepository) Find(promotion models.Promotion) (models.Promotion, error) {
	var promotions models.Promotion
	err := c.db.DB.
		Debug().
		Model(&models.Promotion{}).
		Where(&promotion).
		Take(&promotions).Error
	return promotions, err
}

func (c PromotionRepository) Create(promotion models.Promotion) error {
	return c.db.DB.Create(&promotion).Error
}

func (c PromotionRepository) Update(promotion models.Promotion) error {
	return c.db.DB.Save(&promotion).Error
}

func (c PromotionRepository) Delete(promotion models.Promotion) error {
	return c.db.DB.Delete(&promotion).Error
}
