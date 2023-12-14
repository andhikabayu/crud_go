package models

import "time"

// User represents a user model.
type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
