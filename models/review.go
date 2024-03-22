package models

import "time"

type Review struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"review_id"`
	StoreID     int64     `json:"store_id"`
	UserID      int64     `json:"user_id"`
	Comment     string    `gorm:"type:VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"review_comment"`
	RatingScore int64     `json:"rating_score"`
	RatingStatus bool `json:"rating_status"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Store     Store     `gorm:"foreignKey:StoreID"`
	User        User      `gorm:"foreignKey:UserID"`
}

func (review *Review) TableName() string {
	return "review"
}

func (review *Review) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["review_id"] = review.ID
	resp["store_id"] = review.StoreID
	resp["user_id"] = review.UserID
	resp["username"] = review.User.Username
	resp["review_comment"] = review.Comment
	resp["rating_score"] = review.RatingScore
	resp["rating_status"] = review.RatingStatus
	resp["created_at"] = review.CreatedAt
	resp["updated_at"] = review.UpdatedAt
	return resp
}
