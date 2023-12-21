package repository

import (
	"TableTru/infrastructure"
	"TableTru/models"
)

type PromotionCodeRepository struct {
	db infrastructure.Database
}

func NewPromotionCodeRepository(db infrastructure.Database) PromotionCodeRepository {
	return PromotionCodeRepository{
		db: db,
	}
}

func (c PromotionCodeRepository) FindAll(promotionCode models.PromotionCode, keyword string) (*[]models.PromotionCode, int64, error) {
	var promotionCodes []models.PromotionCode
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.PromotionCode{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			c.db.DB.Where("promotionCode.promotion_code_id LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(promotionCode).
		Find(&promotionCodes).
		Count(&totalRows).Error
	return &promotionCodes, totalRows, err
}

func (c PromotionCodeRepository) Find(promotionCode models.PromotionCode) (models.PromotionCode, error) {
	var promotionCodes models.PromotionCode
	err := c.db.DB.
		Debug().
		Model(&models.PromotionCode{}).
		Where(&promotionCode).
		Take(&promotionCodes).Error
	return promotionCodes, err
}

func (c PromotionCodeRepository) Create(promotionCode models.PromotionCode) error {
	return c.db.DB.Create(&promotionCode).Error
}

func (c PromotionCodeRepository) Update(promotionCode models.PromotionCode) error {
	return c.db.DB.Save(&promotionCode).Error
}

func (c PromotionCodeRepository) Delete(promotionCode models.PromotionCode) error {
	return c.db.DB.Delete(&promotionCode).Error
}
