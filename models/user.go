package models

import "time"

type User struct {
	ID           int64     `gorm:"primary_key;auto_increment" json:"user_id"`
	StoreID      int64     `gorm:"default:null" json:"store_id"`
	Username     string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"username"`
	Password     string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"password"`
	Status       string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"user_status"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Email        string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"email"`
	PhoneNumber  string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"phone_num"`
	ProfileImage string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"profile_image"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	// Store        Store     `gorm:"foreignKey:StoreID"`
}

func (user *User) TableName() string {
	return "user"
}

func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["user_id"] = user.ID
	resp["store_id"] = user.StoreID
	resp["username"] = user.Username
	resp["password"] = user.Password
	resp["user_status"] = user.Status
	resp["latitude"] = user.Latitude
	resp["longitude"] = user.Longitude
	resp["email"] = user.Email
	resp["phone_num"] = user.PhoneNumber
	resp["profile_image"] = user.ProfileImage
	resp["created_at"] = user.CreatedAt
	resp["updated_at"] = user.UpdatedAt
	return resp
}
