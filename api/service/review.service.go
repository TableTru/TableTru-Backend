package service

import (
	"TableTru/api/repository"
	"TableTru/models"
)

type ReviewService struct {
	repository repository.ReviewRepository
}

func NewReviewService(r repository.ReviewRepository) ReviewService {
	return ReviewService{
		repository: r,
	}
}

func (c ReviewService) FindAllReview(review models.Review, keyword string) (*[]models.Review, int64, error) {
	return c.repository.FindAll(review, keyword)
}

func (c ReviewService) FindReview(review models.Review) (models.Review, error) {
	return c.repository.Find(review)
}

func (c ReviewService) CreateReview(review models.Review) error {
	return c.repository.Create(review)
}

func (c ReviewService) UpdateReview(review models.Review) error {
	return c.repository.Update(review)
}

func (c ReviewService) DeleteReview(id int64) error {
	var review models.Review
	review.ID = id
	return c.repository.Delete(review)
}
