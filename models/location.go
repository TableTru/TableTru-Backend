package models

import "time"

type Location struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"location_id"`
	StoreID   int64     `json:"store_id"`
	Name      string    `gorm:"varchar(255)" json:"location_name"`
	ImageName string    `gorm:"varchar(255)" json:"location_image_name"`
	Latitude  string    `gorm:"varchar(255)" json:"latitude"`
	Longitude string    `gorm:"varchar(255)" json:"longitude"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Store     Store     `gorm:"foreignKey:StoreID"`
}

func (location *Location) TableName() string {
	return "location"
}

func (location *Location) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["location_id"] = location.ID
	resp["store_id"] = location.StoreID
	resp["location_name"] = location.ImageName
	resp["location_image_name"] = location.ImageName
	resp["latitude"] = location.ImageName
	resp["longitude"] = location.ImageName
	resp["created_at"] = location.CreatedAt
	resp["updated_at"] = location.UpdatedAt
	return resp
}
