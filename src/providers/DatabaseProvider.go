package providers

import (
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func GetDatabase() *gorm.DB {
	dbOnce.Do(func() {
		var err error

		dsn := os.Getenv("DATABASE_URI")
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("failed to connect to database")
		}

		sqlDb, err := db.DB()

		if err != nil {
			panic("failed to connect to database")
		}

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDb.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDb.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDb.SetConnMaxLifetime(time.Hour)

		if err != nil {
			panic("failed to connect to database")
		}
	})

	return db
}
