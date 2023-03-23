package models

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" gorm:"unique" validate:"email"`
	Password string `json:"password" validate:"required,min=3,max=10"`
	Admin    bool   `json:"admin" validate:"required"`
	Menus    []Menu
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var validate = validator.New()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(schema interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(schema)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
