package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/jainam240101/zomato-clone/scripts/schemas"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GenerateUsers(DB *gorm.DB) {
	for i := 0; i < 20; i++ {
		hashedPassword, _ := HashPassword(faker.Password())
		user := schemas.UserModel{
			ID:          uuid.New(),
			Name:        faker.FirstName() + " " + faker.LastName(),
			Email:       faker.Email(),
			Password:    hashedPassword,
			PhoneNumber: faker.Phonenumber(),
		}
		if err := DB.Save(&user).Error; err != nil {
			fmt.Println("Error ", err)
			return
		}
		fmt.Println("user ---> ", user)
	}

}

func GenerateRestaurants(DB *gorm.DB) {
	for i := 0; i < 20; i++ {
		hashedPassword, _ := HashPassword(faker.Password())

		restro := schemas.RestaurantModel{
			ID:            uuid.New(),
			Name:          faker.FirstName() + " " + faker.LastName(),
			Email:         faker.Email(),
			Password:      hashedPassword,
			PhoneNumber:   faker.Phonenumber(),
			Manager_Name:  faker.FirstName() + " " + faker.LastName(),
			Manager_Phone: faker.Phonenumber(),
			Listing_Date:  faker.Date(),
		}
		if err := DB.Save(&restro).Error; err != nil {
			fmt.Println("Error ", err)
			return
		}
	}
}

func GenerateDrivers(DB *gorm.DB) {
	for i := 0; i < 20; i++ {
		hashedPassword, _ := HashPassword(faker.Password())

		driver := schemas.DriverModel{
			ID:             uuid.New(),
			Name:           faker.FirstName() + " " + faker.LastName(),
			Email:          faker.Email(),
			Password:       hashedPassword,
			PhoneNumber:    faker.Phonenumber(),
			Current_Rating: 0,
			Joining_Date:   time.Now(),
		}
		if err := DB.Save(&driver).Error; err != nil {
			fmt.Println("Error ", err)
			return
		}
	}
}

func main() {
	godotenv.Load("../.env")
	dbUrl := fmt.Sprintf("user=%v password=%v dbname=%v port=5432 sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbUrl,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		fmt.Println("Database Error ----- ", err)
		os.Exit(0)
	}
	DB.AutoMigrate(&schemas.UserModel{}, &schemas.DriverModel{}, &schemas.RestaurantModel{})
	GenerateUsers(DB)
	GenerateRestaurants(DB)
	GenerateDrivers(DB)
}
