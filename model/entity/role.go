package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name     string         `json:"name" gorm:"index:,unique,type:varchar(100)"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}