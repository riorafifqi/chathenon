package entity

import "time"

type User struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time
}
