package models

// Model User
type User struct {
	Username string `json:"username" example:"lala"`
	Password string `json:"password" example:"P@ssw0rd!23"`
	Email    string `json:"email" example:"lala@gmailcom"`
	Name     string `json:"name" example:"lala lili"`
	Birthday string `json:"birthday" example:"1996-04-14"`
}

// Model Login
type Login struct {
	Username string `json:"username" example:"lala"`
	Password string `json:"password" example:"P@ssw0rd!23"`
}
