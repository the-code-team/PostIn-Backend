package providers

import (
	"epsa.upv.es/postin_backend/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
	err    error
)

func GetDatabase() *gorm.DB {
	dbOnce.Do(func() {
		dsn := os.Getenv("DATABASE_URI")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func InitDatabase() {
	db := GetDatabase()

	err := db.AutoMigrate(
		&models.Event{}, &models.Message{},
		&models.Profile{}, &models.Propose{},
	)

	if err != nil {
		panic("failed to migrate database")
	}
}
