package db

import (
	"errors"
	"fmt"
	"strconv"

	"notgithub.com/hyperinactive/api-gateway/config"
	"notgithub.com/hyperinactive/api-gateway/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// conenct to the database
// TODO: error package or smth
func Connect() error {
	var err error
	p := config.Config.Db.Port
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		return errors.New("failed to parse database port")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config.Db.Host,
		port,
		config.Config.Db.User,
		config.Config.Db.Password,
		config.Config.Db.Name,
	)
	// TODO: migration
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return errors.New("failed to connect database")
	}

	DB = database

	fmt.Println("Connection Opened to Database")
	// TODO: migration
	DB.AutoMigrate(&model.User{})

	// DB.AutoMigrate(&User{})
	return nil
}
