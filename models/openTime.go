package models

import "time"

type OpenTime struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"openTime_id"`
	StoreID   int64     `json:"store_id"`
	Day       string    `gorm:"type:VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"day"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (openTime *OpenTime) TableName() string {
	return "openTime"
}

func (openTime *OpenTime) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["openTime_id"] = openTime.ID
	resp["store_id"] = openTime.StoreID
	resp["day"] = openTime.Day
	resp["start_time"] = openTime.StartTime
	resp["end_time"] = openTime.EndTime
	resp["created_at"] = openTime.CreatedAt
	resp["updated_at"] = openTime.UpdatedAt
	return resp
}
