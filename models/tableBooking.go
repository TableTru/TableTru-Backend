package models

import "time"

type TableBooking struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"table_booking_id"`
	StoreID     int64     `json:"store_id"`
	UserID      int64     `json:"user_id"`
	PromotionID int64     `gorm:"default:null" json:"promotion_id"`
	Status      string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"table_booking_status"`
	Count       int64     `json:"table_booking_count"`
	BookingTime time.Time `gorm:"type:datetime;default:NULL" json:"table_booking_time"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Store       Store     `gorm:"foreignKey:StoreID"`
	User        User      `gorm:"foreignKey:UserID"`
	Promotion	Promotion `gorm:"foreignKey:PromotionID"`
}

func (tableBooking *TableBooking) TableName() string {
	return "tableBooking"
}

func (tableBooking *TableBooking) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["table_booking_id"] = tableBooking.ID
	resp["store_id"] = tableBooking.StoreID
	resp["store"] = tableBooking.Store
	resp["user_id"] = tableBooking.UserID
	resp["promotion_id"] = tableBooking.PromotionID
	resp["table_booking_status"] = tableBooking.Status
	resp["table_booking_count"] = tableBooking.Count
	resp["table_booking_time"] = tableBooking.BookingTime
	resp["created_at"] = tableBooking.CreatedAt
	resp["updated_at"] = tableBooking.UpdatedAt
	return resp
}
