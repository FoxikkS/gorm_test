package models

import "gorm.io/gorm"
import "github.com/go-playground/validator/v10"

type User struct {
	gorm.Model
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=6"`
}

var Validate = validator.New()
