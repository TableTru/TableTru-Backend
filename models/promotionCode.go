package models

import "time"

type PromotionCode struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"promotion_code_id"`
	PromotionID int64     `json:"promotion_id"`
	UserID      int64     `json:"user_id"`
	Status      string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"code_status"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Promotion   Promotion `gorm:"foreignKey:PromotionID"`
	User        User      `gorm:"foreignKey:UserID"`
}

func (promotionCode *PromotionCode) TableName() string {
	return "promotionCode"
}

func (promotionCode *PromotionCode) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["promotion_code_id"] = promotionCode.ID
	resp["promotion_id"] = promotionCode.Promotion
	resp["user_id"] = promotionCode.UserID
	resp["code_status"] = promotionCode.Status
	resp["created_at"] = promotionCode.CreatedAt
	resp["updated_at"] = promotionCode.UpdatedAt
	return resp
}
