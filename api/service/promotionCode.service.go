package service

import (
	"TableTru/api/repository"
	"TableTru/models"
)

type PromotionCodeService struct {
	repository repository.PromotionCodeRepository
}

func NewPromotionCodeService(r repository.PromotionCodeRepository) PromotionCodeService {
	return PromotionCodeService{
		repository: r,
	}
}

func (c PromotionCodeService) FindAllPromotionCode(promotionCode models.PromotionCode, keyword string) (*[]models.PromotionCode, int64, error) {
	return c.repository.FindAll(promotionCode, keyword)
}

func (c PromotionCodeService) FindPromotionCode(promotionCode models.PromotionCode) (models.PromotionCode, error) {
	return c.repository.Find(promotionCode)
}

func (c PromotionCodeService) CreatePromotionCode(promotionCode models.PromotionCode) error {
	return c.repository.Create(promotionCode)
}

func (c PromotionCodeService) UpdatePromotionCode(promotionCode models.PromotionCode) error {
	return c.repository.Update(promotionCode)
}

func (c PromotionCodeService) DeletePromotionCode(id int64) error {
	var promotionCode models.PromotionCode
	promotionCode.ID = id
	return c.repository.Delete(promotionCode)
}
