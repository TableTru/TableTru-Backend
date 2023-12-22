package service

import (
	"TableTru/api/repository"
	"TableTru/models"
)

type RatingService struct {
	repository repository.RatingRepository
}

func NewRatingService(r repository.RatingRepository) RatingService {
	return RatingService{
		repository: r,
	}
}

func (c RatingService) FindAllRating(rating models.Rating, keyword string) (*[]models.Rating, int64, error) {
	return c.repository.FindAll(rating, keyword)
}

func (c RatingService) FindRating(rating models.Rating) (models.Rating, error) {
	return c.repository.Find(rating)
}

func (c RatingService) CreateRating(rating models.Rating) error {
	return c.repository.Create(rating)
}

func (c RatingService) UpdateRating(rating models.Rating) error {
	return c.repository.Update(rating)
}

func (c RatingService) DeleteRating(id int64) error {
	var rating models.Rating
	rating.ID = id
	return c.repository.Delete(rating)
}
