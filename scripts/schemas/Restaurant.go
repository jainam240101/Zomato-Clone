package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RestaurantModel struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:char(36);primary_key"`
	Name          string    `json:"name"`
	Email         string    `json:"email" gorm:"unique"`
	Password      string    `json:"password"`
	PhoneNumber   string    `json:"phoneNumber" gorm:"unique"`
	Manager_Name  string    `json:"managerName"`
	Manager_Phone string    `json:"managerPhone"`
	Listing_Date  string    `json:"listingDate"`
}
