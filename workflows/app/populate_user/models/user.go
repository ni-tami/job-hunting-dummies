package models

type CreateUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
