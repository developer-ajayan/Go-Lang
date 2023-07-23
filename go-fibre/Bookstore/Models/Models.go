package Models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

var Books []Book
