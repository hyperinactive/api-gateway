package handler

import (
	"errors"
	"fmt"

	"notgithub.com/hyperinactive/api-gateway/db"
	"notgithub.com/hyperinactive/api-gateway/model"
)

func GetUserById(id uint64) (*model.User, error) {
	var user model.User

	db.DB.Find(&user, id)

	if user.Username == "" {
		return nil, errors.New(fmt.Sprintf("No user found with ID %v", id))
	}

	return &user, nil
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User

	db.DB.Where("username = ?", username).First(&user)

	// user won't be nil, the object will be there, just empty
	if user.Username == "" {
		return nil, errors.New(fmt.Sprintf("No user found with username %s", username))
	}

	return &user, nil
}

func CreateUser(username string, email string, password string) error {
	hash, err := HashPassword(password)

	if err != nil {
		return err
	}

	user := new(model.User)
	user.Username = username
	user.Password = hash
	user.Email = email

	result := db.DB.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	fmt.Printf("%+v", user)

	return nil
}
