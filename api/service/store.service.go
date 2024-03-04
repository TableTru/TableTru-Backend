package service

import (
	"TableTru/api/repository"
	"TableTru/models"
)

type StoreService struct {
	repository repository.StoreRepository
}

func NewStoreService(r repository.StoreRepository) StoreService {
	return StoreService{
		repository: r,
	}
}

func (c StoreService) FindAllStore(store models.Store, keyword string) (*[]models.Store, int64, error) {
	return c.repository.FindAll(store, keyword)
}

func (c StoreService) FindStore(store models.Store) (models.Store, error) {
	return c.repository.Find(store)
}

func (c StoreService) CreateStore(store models.Store) error {
	return c.repository.Create(store)
}

func (c StoreService) UpdateStore(store models.Store) error {
	return c.repository.Update(store)
}

func (c StoreService) DeleteStore(id int64) error {
	var store models.Store
	store.ID = id
	return c.repository.Delete(store)
}