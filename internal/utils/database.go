package utils

import (
	"errors"
	"github.com/MrRytis/chat-api/pkg/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB

func NewDb() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            false,
	})

	if err != nil {
		log.Fatal(err, "Error connecting to database")
	}

	Db = db

	return db
}

func HandleDbError(err error) {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exception.NewNotFound("Record not found")
		}

		exception.NewInternalServerError()
	}
}
