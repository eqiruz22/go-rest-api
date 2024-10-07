package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"type:varchar(100)"`
	Email     string `json:"email" gorm:"index:,unique,type:varchar(100)"`
	Address   string `json:"address" gorm:"type:varchar(100)"`
	Phone     string `json:"phone" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}