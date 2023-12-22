package service

import (
	"TableTru/api/repository"
	"TableTru/models"
)

type ReviewImageService struct {
	repository repository.ReviewImageRepository
}

func NewReviewImageService(r repository.ReviewImageRepository) ReviewImageService {
	return ReviewImageService{
		repository: r,
	}
}

func (c ReviewImageService) FindAllReviewImage(reviewImage models.ReviewImage, keyword string) (*[]models.ReviewImage, int64, error) {
	return c.repository.FindAll(reviewImage, keyword)
}

func (c ReviewImageService) FindReviewImage(reviewImage models.ReviewImage) (models.ReviewImage, error) {
	return c.repository.Find(reviewImage)
}

func (c ReviewImageService) CreateReviewImage(reviewImage models.ReviewImage) error {
	return c.repository.Create(reviewImage)
}

func (c ReviewImageService) UpdateReviewImage(reviewImage models.ReviewImage) error {
	return c.repository.Update(reviewImage)
}

func (c ReviewImageService) DeleteReviewImage(id int64) error {
	var reviewImage models.ReviewImage
	reviewImage.ID = id
	return c.repository.Delete(reviewImage)
}
