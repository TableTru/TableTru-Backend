package models

import "time"

type User struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"user_id"`
	Username  string    `gorm:"varchar(255)" json:"username"`
	Password  string    `gorm:"varchar(255)" json:"password"`
	Status    string    `gorm:"varchar(255)" json:"user_status"`
	Latitude  string    `gorm:"varchar(255)" json:"latitude"`
	Longitude string    `gorm:"varchar(255)" json:"longitude"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (user *User) TableName() string {
	return "user"
}

func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["user_id"] = user.ID
	resp["username"] = user.Username
    resp["password"] = user.Password
    resp["user_status"] = user.Status
    resp["latitude"] = user.Latitude
    resp["longitude"] = user.Longitude
	resp["created_at"] = user.CreatedAt
	resp["updated_at"] = user.UpdatedAt
	return resp
}
