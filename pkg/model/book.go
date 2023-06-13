package model

import "gorm.io/gorm"

// Book represents a book
//
// swagger:model
type Book struct {
	// the id for this book
	// required: false
	gorm.Model

	// the name of the book
	// required: true
	// min length: 3
	// example: Moondust
	// unique: true
	Title string `json:"title"`

	// the book author's name
	// required: true
	// min length: 3
	// example: Andrew Smith
	Author string `json:"author"`

	// the book's description
	// required: true
	// min length: 3
	// example: test
	Description string `json:"description"`
}
