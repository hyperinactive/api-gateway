package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Init() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	Config = &ConfigStruct{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Db: Db{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
	return nil
}

var Config *ConfigStruct
