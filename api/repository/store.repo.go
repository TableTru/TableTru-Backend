package repository

import (
	"TableTru/infrastructure"
	"TableTru/models"
)

type StoreRepository struct {
	db infrastructure.Database
}

func NewStoreRepository(db infrastructure.Database) StoreRepository {
	return StoreRepository{
		db: db,
	}
}

func (c StoreRepository) FindAll(store models.Store, keyword string) (*[]models.Store, int64, error) {
	var stores []models.Store
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.Store{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			c.db.DB.Where("store.store_name LIKE ? ", queryKeyword))
	}

	err := queryBuider.
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
