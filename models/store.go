package models

import "time"

type Store struct {
	ID               int64      `gorm:"primary_key;auto_increment" json:"store_id"`
	CategoryID       int64      `json:"category_id"`
	Name             string     `gorm:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"store_name"`
	Description      string     `gorm:"type:VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"store_description"`
	CoverImage       string     `gorm:"type:VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"store_cover_image"`
	TableBooking     int64      `json:"table_booking"`
	MaxPeopleBooking int64      `json:"max_people_booking"`
	SumRating        float64    `json:"sum_rating"`
	Latitude         float64    `json:"latitude"`
	Longitude        float64    `json:"longitude"`
	Location        string    `gorm:"type:VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"location"`
	OpenTimes        []OpenTime `gorm:"foreignKey:StoreID"`
	CreatedAt        time.Time  `json:"created_at,omitempty"`
	UpdatedAt        time.Time  `json:"updated_at,omitempty"`
	Category         Category   `gorm:"foreignKey:CategoryID"`
}

func (store *Store) TableName() string {
	return "store"
}

func (store *Store) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["store_id"] = store.ID
	resp["category_id"] = store.CategoryID
	resp["location"] = store.Location
	resp["store_name"] = store.Name
	resp["store_description"] = store.Description
	resp["store_cover_image"] = store.CoverImage
	resp["table_booking"] = store.TableBooking
	resp["max_people_booking"] = store.MaxPeopleBooking
	resp["sum_rating"] = store.SumRating
	resp["latitude"] = store.Latitude
	resp["longitude"] = store.Longitude
	resp["OpenTimes"] = store.OpenTimes
	resp["category_name"] = store.Category.Name
	resp["created_at"] = store.CreatedAt
	resp["updated_at"] = store.UpdatedAt
	return resp
}
