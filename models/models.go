package models

type User struct {
	ID         int64  `json:"id"`
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
