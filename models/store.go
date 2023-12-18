package models

import "time"

type Store struct {
	ID           int64     `gorm:"primary_key;auto_increment" json:"store_id"`
	CategoryID   int64     `json:"category_id"`
	Name         string    `gorm:"varchar(255)" json:"store_name"`
	Description  string    `gorm:"varchar(500)" json:"store_description"`
	ImageName    string    `gorm:"varchar(255)" json:"store_menu_image"`
	TableBooking int       `json:"table_booking"`
	SumRating    float64   `json:"sum_rating"`
	OpenTime     time.Time `json:"open_time"`
	CloseTime     time.Time `json:"close_time"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	Category     Category  `gorm:"foreignKey:CategoryID"`
}

func (store *Store) TableName() string {
    return "store"
}

func (store *Store) ResponseMap() map[string]interface{} {
    resp := make(map[string]interface{})
    resp["store_id"] = store.ID
    resp["category_id"] = store.CategoryID
    resp["store_name"] = store.Name
	resp["store_description"] = store.Description
	resp["store_menu_image"] = store.ImageName
	resp["table_booking"] = store.TableBooking
	resp["sum_rating"] = store.SumRating
	resp["open_time"] = store.OpenTime
	resp["close_time"] = store.CloseTime
    resp["created_at"] = store.CreatedAt
    resp["updated_at"] = store.UpdatedAt
    return resp
}