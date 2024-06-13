package model

type Book struct {
	// gorm.Model
	Base
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
