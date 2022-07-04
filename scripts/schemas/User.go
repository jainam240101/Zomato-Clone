package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:char(36);primary_key"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phoneNumber" gorm:"unique"`
}
