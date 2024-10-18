package response

import (
	"fiber/backend/model/entity"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	RoleId    int 		`json:"role_id"`
	Role      entity.Role `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}