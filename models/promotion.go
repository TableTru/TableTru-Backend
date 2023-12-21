package models

import "time"

type Promotion struct {
	ID             int64     `gorm:"primary_key;auto_increment" json:"promotion_id"`
	StoreID        int64     `json:"store_id"`
	Name           string    `gorm:"varchar(255)" json:"promotion_name"`
	Description    string    `gorm:"varchar(255)" json:"promotion_description"`
	ExpirationDate time.Time `gorm:"type:datetime;default:NULL" json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	Store          Store     `gorm:"foreignKey:StoreID"`
}

func (promotion *Promotion) TableName() string {
	return "promotion"
}

func (promotion *Promotion) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["promotion_id"] = promotion.ID
	resp["store_id"] = promotion.StoreID
	resp["promotion_name"] = promotion.Name
	resp["promotion_description"] = promotion.Description
	resp["expiration_date"] = promotion.ExpirationDate
	resp["created_at"] = promotion.CreatedAt
	resp["updated_at"] = promotion.UpdatedAt
	return resp
}
