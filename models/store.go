package models

import "time"

type Store struct {
	ID           int64      `gorm:"primary_key;auto_increment" json:"store_id"`
	CategoryID   int64      `json:"category_id"`
	LocationID   int64      `json:"location_id"`
	Name         string     `gorm:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"store_name"`
	Description  string     `gorm:"type:VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"store_description"`
	TableBooking int        `json:"table_booking"`
	SumRating    float64    `json:"sum_rating"`
	Latitude     string     `gorm:"varchar(255)" json:"latitude"`
	Longitude    string     `gorm:"varchar(255)" json:"longitude"`
	OpenTimes    []OpenTime `gorm:"foreignKey:StoreID"`
	CreatedAt    time.Time  `json:"created_at,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty"`
	Category     Category   `gorm:"foreignKey:CategoryID"`
	Location     Location   `gorm:"foreignKey:LocationID"`
}


func (store *Store) TableName() string {
	return "store"
}

func (store *Store) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["store_id"] = store.ID
	resp["category_id"] = store.CategoryID
	resp["location_id"] = store.LocationID
	resp["store_name"] = store.Name
	resp["store_description"] = store.Description
	resp["table_booking"] = store.TableBooking
	resp["sum_rating"] = store.SumRating
	resp["latitude"] = store.Latitude
	resp["longitude"] = store.Longitude
	resp["OpenTimes"] = store.OpenTimes
	resp["created_at"] = store.CreatedAt
	resp["updated_at"] = store.UpdatedAt
	return resp
}
