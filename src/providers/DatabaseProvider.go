package providers

import (
	"sync"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	dbOnce sync.Once
)

func GetDatabase() *gorm.DB {
	dbOnce.Do(func() {
		var err error

		dsn := os.Getenv("DATABASE_URI")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("failed to connect to database")
		}
	})

	return db
}