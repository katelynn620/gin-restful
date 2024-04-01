package database

import (
	"github.com/katelynn620/gin-restful/pkg/model"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

type DatabaseManager struct {
	DB *gorm.DB
}

func InitDatabaseManager() (*DatabaseManager, error) {
	logger := zap.L().Sugar()
	defer logger.Sync()

	var (
		db  *gorm.DB
		err error
	)
	url := viper.GetString("database.url")
	gormlogger := zapgorm2.New(zap.L())
	gormlogger.SetAsDefault()
	db, err = gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: gormlogger.LogMode(glogger.Info),
	})
	if err != nil {
		logger.Errorln("failed to connect database")
		return nil, err
	}
	return &DatabaseManager{
		DB: db,
	}, nil
}

func (d *DatabaseManager) Migrate() (err error) {
	logger := zap.L().Sugar()
	defer logger.Sync()

	logger.Debug("Migrating database")
	err = d.DB.AutoMigrate(&model.Book{})
	if err != nil {
		logger.Panicf("failed to migrate database: %v", err)
	}
	return
}
