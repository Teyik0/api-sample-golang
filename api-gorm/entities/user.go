package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"name"`
	Password string `json:"breed"`
}
