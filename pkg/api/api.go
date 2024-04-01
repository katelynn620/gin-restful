package api

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/katelynn620/gin-restful/pkg/config"
	database "github.com/katelynn620/gin-restful/pkg/db"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func Init() (r *gin.Engine) {
	logger := zap.L().Sugar()
	defer logger.Sync()

	_ = config.GetConfig()

	dbm, err := database.InitDatabaseManager()
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}

	dbm.Migrate()
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}

	h := handler{
		DB: dbm.DB,
	}

	debug := viper.GetBool("debug")
	logger.Debugf("Debug mode :%b", debug)
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r = gin.Default()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	r.SetTrustedProxies(nil)

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
