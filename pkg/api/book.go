package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/katelynn620/gin-restful/pkg/model"
)

type BookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) GetBooks(ctx *gin.Context) {
	var books []model.Book

	if r := h.DB.Find(&books); r.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, r.Error)
		return
	}

	ctx.JSON(http.StatusOK, &books)
}

func (h handler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book model.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}

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
