package db

type User struct {
	Model
	Username string
	Password string
	Salt     string
}
