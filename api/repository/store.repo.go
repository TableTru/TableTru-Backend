package repository

import (
	"TableTru/infrastructure"
	"TableTru/models"

	"context"
	"fmt"
	"log"

	"googlemaps.github.io/maps"
)

type StoreRepository struct {
	db infrastructure.Database
}

func NewStoreRepository(db infrastructure.Database) StoreRepository {
	return StoreRepository{
		db: db,
	}
}

func (r *StoreRepository) SeedData(stores []models.Store) error {
	result := r.db.DB.Create(stores)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c StoreRepository) FindAll(store models.Store, keyword string) (*[]models.Store, int64, error) {
	var stores []models.Store
	var totalRows int64 = 0

	queryBuilder := c.db.DB.Order("created_at desc").Model(&models.Store{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where("store_name LIKE ?", queryKeyword)
	}

	err := queryBuilder.
		Preload("Category").
		Preload("OpenTimes").
		Where(store).
		Find(&stores).
		Count(&totalRows).Error
	return &stores, totalRows, err
}

func (c StoreRepository) Find(store models.Store) (models.Store, error) {
	var stores models.Store
	err := c.db.DB.
		Debug().
		Model(&models.Store{}).
		Preload("Category").
		Preload("OpenTimes").
		Where(&store).
		Take(&stores).Error
	return stores, err
}

func (c StoreRepository) Create(store models.Store) error {
	return c.db.DB.Create(&store).Error
}

func (c StoreRepository) Update(store models.Store) error {
	return c.db.DB.Save(&store).Error
}

func (c StoreRepository) Delete(store models.Store) error {
	return c.db.DB.Delete(&store).Error
}

func (c StoreRepository) FindbyNumber(store models.Store, keyword string, num int) (*[]models.Store, int64, error) {
	var stores []models.Store
	var totalRows int64 = 0

	queryBuilder := c.db.DB.Order("created_at desc").Model(&models.Store{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where("store_name LIKE ?", queryKeyword)
	}

	err := queryBuilder.
		Limit(num).
		Preload("OpenTimes").
		Where(store).
		Find(&stores).
		Count(&totalRows).Error
	return &stores, totalRows, err
}

func (c StoreRepository) SearchStoreRatingSort(store models.Store, keyword string) (*[]models.Store, int64, error) {
	var stores []models.Store
	var totalRows int64 = 0

	queryBuilder := c.db.DB.Order("sum_rating desc").Model(&models.Store{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where("store_name LIKE ?", queryKeyword)
	}

	err := queryBuilder.
		Preload("Category").
		Preload("OpenTimes").
		Where(store).
		Find(&stores).
		Count(&totalRows).Error
	return &stores, totalRows, err
}

func (c StoreRepository) SearchStoreLocationSort(store models.Store, originLocation string, keyword string) (*[]models.Store, int64, error) {
	apiKey := "AIzaSyAuPMpFFFGWcocgnd1axgKUHSa2poG9rNY"
	client, apiErr := maps.NewClient(maps.WithAPIKey(apiKey))
	if apiErr != nil {
		log.Fatalf("Failed to create client: %v", apiErr)
	}

	var stores []models.Store
	var totalRows int64 = 0

	queryBuilder := c.db.DB.Order("sum_rating desc").Model(&models.Store{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where("store_name LIKE ?", queryKeyword)
	}

	err := queryBuilder.
		Preload("Category").
		Preload("OpenTimes").
		Where(store).
		Find(&stores).
		Count(&totalRows).Error

	destinations := []string{"Chiang Mai, Thailand", "Phuket, Thailand", "Pattaya, Thailand"}

	r := &maps.DistanceMatrixRequest{
		Origins:      []string{originLocation},
		Destinations: destinations,
	}

	resp, err := client.DistanceMatrix(context.Background(), r)
	if err != nil {
		log.Fatalf("DistanceMatrix request failed: %v", err)
	}

	return &stores, totalRows, err
}
