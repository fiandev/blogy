package database

import (
	"blogy/config"
	"blogy/models"
	"fmt"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USERNAME"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(
		&models.Post{},
		&models.User{},
	)
	fmt.Println("Database Migrated")
}
