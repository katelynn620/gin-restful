package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/katelynn620/gin-restful/pkg/model"
)

// Book model in request body
//
// swagger:model
type BookRequestBody struct {
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

// swagger:operation GET /books Books listBooks
//
// Returns list of books
// ---
//
//	Produces:
//	- application/json
//
//	Responses:
//		200:
//			description: Successful operation
//			schema:
//				type: array
//				items:
//					"$ref": "#/definitions/Book"
//		500:
//			description: internal server error
func (h handler) GetBooks(ctx *gin.Context) {
	var books []model.Book

	if r := h.DB.Find(&books); r.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, r.Error)
		return
	}

	ctx.JSON(http.StatusOK, &books)
}

// swagger:operation GET /books/{id} Books getBook
// Returns a book
// ---
//
//	Parameters:
//	- name: id
//		in: path
//		description: ID of the book
//		required: true
//		type: string
//	Produces:
//	- application/json
//	Responses:
//		'200':
//			description: Successful operation
//			schema:
//				items:
//					"$ref": "#/definitions/Book"
//		'404':
//			description: 'Error: Not Found'
func (h handler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book model.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}

// swagger:operation POST /books Books addBook
//
// Adds a new book
// ---
//
//	Consumes:
//	- application/json
//	Produces:
//	- application/json
//	Parameters:
//	- name: book
//		in: body
//		description: The new book to create
//		schema:
//			"$ref": "#/definitions/BookRequestBody"
//	Responses:
//		'201':
//			description: Successful operation
//			schema:
//				type: array
//				items:
//					"$ref": "#/definitions/Book"
//		'400':
//			description: invalid input
//		'500':
//			description: internal server error
func (h handler) AddBook(ctx *gin.Context) {
	body := BookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book model.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}

// swagger:operation PUT /books/{id} Books updateBook
//
// Update an existing book
// ---
//
//	Parameters:
//	- name: id
//		in: path
//		description: ID of the book
//		required: true
//		type: string
//	- name: book
//		in: body
//		description: The book to update
//		schema:
//			"$ref": "#/definitions/BookRequestBody"
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Responses:
//			'200':
//				description: Successful operation
//				schema:
//					type: array
//					items:
//						"$ref": "#/definitions/Book"
//			'400':
//				description: Invalid input
//			'404':
//				description: book not found
//			'500':
//				description: internal server error
func (h handler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	body := BookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book model.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if body.Title != "" {
		book.Title = body.Title
	}
	if body.Author != "" {
		book.Author = body.Author
	}
	if body.Description != "" {
		book.Description = body.Description
	}

	h.DB.Save(&book)

	ctx.JSON(http.StatusOK, &book)
}

// swagger:operation DELETE /books/{id} Books deleteBook
// Delete an existing book
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the book
//     required: true
//     type: string
//
// consumes:
//   - application/json
//
// produces:
// - application/json
//
//	Responses:
//		'204':
//			description: "Book has been deleted"
//		'404':
//			description: 'Error: Not Found'
func (h handler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book model.Book
	if r := h.DB.First(&book, id); r.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, r.Error)
		return
	}

	h.DB.Delete(&book)

	ctx.Status(http.StatusNoContent)
}
