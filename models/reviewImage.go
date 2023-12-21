package models

import "time"

type ReviewImage struct {
	ID           int64     `gorm:"primary_key;auto_increment" json:"review_image_id"`
	ReviewID   int64     `json:"review_id"`
	ImageName         string    `gorm:"varchar(255)" json:"review_image_name"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	Review     Review  `gorm:"foreignKey:ReviewID"`
}


func (reviewImage *ReviewImage) TableName() string {
	return "reviewImage"
}

func (reviewImage *ReviewImage) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["review_image_id"] = reviewImage.ID
	resp["review_id"] = reviewImage.ReviewID
	resp["review_image_name"] = reviewImage.ImageName
	resp["created_at"] = reviewImage.CreatedAt
	resp["updated_at"] = reviewImage.UpdatedAt
	return resp
}
