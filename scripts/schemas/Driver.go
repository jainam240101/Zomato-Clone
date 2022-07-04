package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DriverModel struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:char(36);primary_key"`
	Name           string    `json:"name"`
	Email          string    `json:"email" gorm:"unique"`
	Password       string    `json:"password"`
	PhoneNumber    string    `json:"phoneNumber" gorm:"unique"`
	Current_Rating float64   `json:"currentRating"`
	Joining_Date   time.Time `json:"start_date"`
}
