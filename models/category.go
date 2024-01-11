package models

import "time"

type Category struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"category_id"`
	Name      string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"category_name"`
	ImageName string    `gorm:"varchar(255)" json:"category_image"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (category *Category) TableName() string {
    return "category"
}

func (category *Category) ResponseMap() map[string]interface{} {
    resp := make(map[string]interface{})
    resp["category_id"] = category.ID
    resp["category_name"] = category.Name
    resp["category_image"] = category.ImageName
    resp["created_at"] = category.CreatedAt
    resp["updated_at"] = category.UpdatedAt
    return resp
}