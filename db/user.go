package db

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	Model
	Username string
	Password string
	Salt     string
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func CalculatePassword(password, salt string) string {
	return strings.ToLower(md5V(fmt.Sprintf("%s:%s", password, salt)))
}
func CheckUser(username string) bool {
	err := Connection.Where("username  = ?", username).Row().Err()
	return err == gorm.ErrRecordNotFound
}
func CreateUser(username, password string) error {
	return Connection.Create(&User{Username: username, Password: password}).Error
}
