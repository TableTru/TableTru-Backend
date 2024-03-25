package service

import (
	"TableTru/api/repository"
	"TableTru/models"
	"fmt"
	"sort"
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

func (c StoreService) FindStoreByNum(store models.Store, keyword string, num int) (*[]models.Store, int64, error) {
	return c.repository.FindbyNumber(store, keyword, num)
}

func (c StoreService) SearchStoreRatingSort(store models.Store, keyword string) (*[]models.Store, int64, error) {
	return c.repository.SearchStoreRatingSort(store, keyword)
}

func (c StoreService) SearchStoreLocationSort(store models.Store, originLocation string, keyword string) (*[]models.Store, int64, error) {
    stores, totalRows, distancesArray, err := c.repository.SearchStoreLocationSort(store, originLocation, keyword)
	sort.Slice(distancesArray, func(i, j int) bool {
		return distancesArray[i].Distance < distancesArray[j].Distance
	})

	// Print the sorted distances
	for _, distance := range distancesArray {
		fmt.Printf("Store ID: %d, Distance: %d\n", distance.StoreID, distance.Distance)
	}

    return stores, totalRows, err
}