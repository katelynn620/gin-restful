package db

import (
	"github.com/katelynn620/gin-restful/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(url string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return
	}

	db.AutoMigrate(&model.Book{})

	return
}
