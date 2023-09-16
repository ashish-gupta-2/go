package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	utils "ashish.com/m/internal/utils"
	"ashish.com/m/pkg/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenDatabase connects to postgres db and returns the database instance.
func OpenDatabase(ctx context.Context, connString string) (*gorm.DB, error) {
	maxRetries := 5
	retryLength := time.Nanosecond
	var db *gorm.DB
	var err error
	for i := 1; i <= maxRetries; i++ {
		// wait and sleep
		select {
		case <-ctx.Done():
			return nil, errors.New("context cancelled")
		case <-time.After(retryLength):
			if retryLength == time.Nanosecond {
				retryLength = utils.WaitTiny
			} else {
				retryLength *= 2
			}
		}

		// open database connection
		db, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
		if db != nil && err == nil {
			connCount := sqlDB(db).Stats().OpenConnections
			fields := log.Fields{"connections": connCount, "attempt": i}
			log.WithContext(ctx).WithFields(fields).Info("db connection established")
			break
		}
		log.WithContext(ctx).WithField("attempt", i).Errorf("db connection error: %v", err)
	}
	if err != nil {
		msg := fmt.Sprintf("db connection failed after multiple attempts: %v", err)
		log.WithContext(ctx).Error(msg)
		return nil, errors.New(msg)
	}

	// set database params
	sqlDB(db).SetMaxIdleConns(10)
	sqlDB(db).SetMaxOpenConns(10)

	// set db migration
	_ = db.AutoMigrate(&models.Employee{})
	_ = db.AutoMigrate(&models.Person{})

	log.WithContext(ctx).Info("data migration done")

	// return the database
	return db, nil
}

func sqlDB(db *gorm.DB) *sql.DB {
	d, _ := db.DB()
	return d
}
