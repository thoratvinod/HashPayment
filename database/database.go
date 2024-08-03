package database

import (
	"fmt"
	"log"

	"github.com/thoratvinod/HashPayment/specs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDatabase() error {
	host := "localhost"
	port := 5432
	dbName := "HashPayment"
	dbUser := "postgres"
	password := ""
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, dbUser, dbName, password)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("database connection failed, err=%+v", err.Error())
		return err
	}
	err = migrate()
	if err != nil {
		return fmt.Errorf("database migration failed, err=%+v", err.Error())
	}
	return nil
}

func migrate() error {
	err := DB.AutoMigrate(specs.PaymentModel{})
	if err != nil {
		return fmt.Errorf("database migration failed, err=%+v", err.Error())
	}
	return nil
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get SQL DB object:", err)
	}
	sqlDB.Close()
}
