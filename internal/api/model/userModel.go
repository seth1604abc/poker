package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Name        string `gorm:"type:varchar(100);not null;column:name"`
	Account     string `gorm:"type:varchar(100);not null;column:account"`
	Password    string `gorm:"type:varchar(256);not null;column:password"`
	Email       string `gorm:"type:varchar(100);not null;column:email"`
	IsActivated bool   `gorm:"type:boolean;default:false;column:is_activated"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (User) TableName() string {
	return "user"
}