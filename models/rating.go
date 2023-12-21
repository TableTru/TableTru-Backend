package models

import "time"

type Rating struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"rating_id"`
	StoreID     int64     `json:"store_id"`
	UserID      int64     `json:"user_id"`
	Score       int64     `json:"rating_score"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Store       Store     `gorm:"foreignKey:StoreID"`
	User        User      `gorm:"foreignKey:UserID"`
}

func (rating *Rating) TableName() string {
	return "rating"
}

func (rating *Rating) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["rating_id"] = rating.ID
	resp["store_id"] = rating.StoreID
	resp["user_id"] = rating.UserID
	resp["rating_score"] = rating.Score
	resp["created_at"] = rating.CreatedAt
	resp["updated_at"] = rating.UpdatedAt
	return resp
}
