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

func (c StoreService) SearchStoreLocationSort(store models.Store, originLocation string, keyword string) ([]models.Store, int64, error) {
    stores, totalRows, distancesArray, err := c.repository.SearchStoreLocationSort(store, originLocation, keyword)

	filter := func(arr []models.StoreDistanceWithIndex) []models.StoreDistanceWithIndex {
		var filtered []models.StoreDistanceWithIndex
		for _, d := range arr {
			if d.StoreLocationStatus {
				filtered = append(filtered, d)
			}
		}
		return filtered
	}

	// Filter out elements where StoreLocationStatus is false
	newDistances := filter(distancesArray)
	
	sort.Slice(newDistances, func(i, j int) bool {
		return newDistances[i].Distance < newDistances[j].Distance
	})

	// Print the sorted distances
	for _, distance := range newDistances {
		fmt.Printf("Store ID: %d, Distance: %d\n", distance.StoreID, distance.Distance)
	}

	filterStores := func(arr []models.Store, distArr []models.StoreDistanceWithIndex) []models.Store {
        var filtered []models.Store
        for _, s := range arr {
            for _, d := range distArr {
                if s.ID == d.StoreID {
                    filtered = append(filtered, s)
                    break
                }
            }
        }
        return filtered
    }

    // Filter stores based on ID matching StoreID in newDistances
    filteredStores := filterStores(*stores, newDistances)

	// Create a map of StoreID to index in newDistances
    distanceIndexMap := make(map[int64]int)
    for i, distance := range newDistances {
        distanceIndexMap[distance.StoreID] = i
    }

    sort.SliceStable(filteredStores, func(i, j int) bool {
		return distanceIndexMap[int64(filteredStores[i].ID)] < distanceIndexMap[int64(filteredStores[j].ID)]
	})


    return filteredStores, totalRows, err
}