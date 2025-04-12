package bootstrap

import (
	"fmt"
	"log"
	"shope_clone/internal/user/domain"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	GormDB *gorm.DB
}

func ConnectDB(cfg *Config) *Database {
	dsn := cfg.DBUrl
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	sqlDB, err := db.DB()
	db.AutoMigrate(domain.User{})
	if err != nil {
		log.Fatalf("Failed to get sqlDB: %v", err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Connected to MySQL with GORM!")
	return &Database{GormDB: db}
}
