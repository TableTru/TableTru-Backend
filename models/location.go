package models

import "time"

type Location struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"location_id"`
	Name      string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"location_name"`
	ImageName string    `gorm:"varchar(255)" json:"location_image_name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (location *Location) TableName() string {
	return "location"
}

func (location *Location) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["location_id"] = location.ID
	resp["location_name"] = location.Name
	resp["location_image_name"] = location.ImageName
	resp["created_at"] = location.CreatedAt
	resp["updated_at"] = location.UpdatedAt
	return resp
}
