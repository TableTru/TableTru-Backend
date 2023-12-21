package models

import "time"

type StoreImage struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"store_image_id"`
	StoreID   int64     `json:"store_id"`
	ImageName string    `gorm:"varchar(255)" json:"store_image_name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Store     Store     `gorm:"foreignKey:StoreID"`
}

func (storeImage *StoreImage) TableName() string {
	return "storeImage"
}

func (storeImage *StoreImage) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["store_image_id"] = storeImage.ID
	resp["store_id"] = storeImage.StoreID
	resp["store_image_name"] = storeImage.ImageName
	resp["created_at"] = storeImage.CreatedAt
	resp["updated_at"] = storeImage.UpdatedAt
	return resp
}