package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint
	Items       []Item
}

type MenuCreateReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type Item struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageURL    string `json:"image_url"`
	MenuID      uint
}

type ItemCreateRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description"`
	Price       int    `json:"price" validate:"required"`
}
