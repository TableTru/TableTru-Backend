package service

import (
	"TableTru/api/repository"
	"TableTru/models"
)

type PromotionService struct {
	repository repository.PromotionRepository
}

func NewPromotionService(r repository.PromotionRepository) PromotionService {
	return PromotionService{
		repository: r,
	}
}

func (c PromotionService) FindAllPromotion(promotion models.Promotion, keyword string) (*[]models.Promotion, int64, error) {
	return c.repository.FindAll(promotion, keyword)
}

func (c PromotionService) FindPromotion(promotion models.Promotion) (models.Promotion, error) {
	return c.repository.Find(promotion)
}

func (c PromotionService) CreatePromotion(promotion models.Promotion) error {
	return c.repository.Create(promotion)
}

func (c PromotionService) UpdatePromotion(promotion models.Promotion) error {
	return c.repository.Update(promotion)
}

func (c PromotionService) DeletePromotion(id int64) error {
	var promotion models.Promotion
	promotion.ID = id
	return c.repository.Delete(promotion)
}
