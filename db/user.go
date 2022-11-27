package db

import "gorm.io/gorm"

type User struct {
	Model
	Username string
	Password string
	Salt     string
}

func CheckUser(username string) bool {
	err := Connection.Where("username  = ?", username).Row().Err()
	return err == gorm.ErrRecordNotFound
}
func CreateUser(username, password string) error {
	return Connection.Create(&User{Username: username, Password: password}).Error
}
