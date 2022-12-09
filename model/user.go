package model

type Register struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"admin"`
}

type Login Register
