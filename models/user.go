package models
import "time"

type User struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
    Name     string    `gorm:"size:200" json:"name"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (user *User) TableName() string {
	return "user"
}

func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
    resp["id"] = user.ID
    resp["name"] = user.Name
    resp["created_at"] = user.CreatedAt
    resp["updated_at"] = user.UpdatedAt
    return resp
}