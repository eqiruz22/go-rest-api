package entity

import (
	"time"

	"gorm.io/gorm"
)

type Auth struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"index:,unique,type:varchar(100)"`
	Password  string 		 `json:"password" gorm:"type:varchar(100)"`
	Role string `json:"role" gorm:"type:varchar(100)"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}

