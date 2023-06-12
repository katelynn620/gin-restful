package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/katelynn620/gin-restful/pkg/config"
	"github.com/katelynn620/gin-restful/pkg/db"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func Init() (r *gin.Engine) {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	dbHandler, err := db.InitDB(c.DBUrl)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	h := handler{
		DB: dbHandler,
	}

	r = gin.Default()

	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	v1 := r.Group("/v1")

	v1.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "helloworld"})
	})

	// books
	books := v1.Group("/books")
	books.POST("/", h.AddBook)
	books.GET("/", h.GetBooks)
	books.GET("/:id", h.GetBook)
	books.PUT("/:id", h.UpdateBook)
	books.DELETE("/:id", h.DeleteBook)

	return
}
