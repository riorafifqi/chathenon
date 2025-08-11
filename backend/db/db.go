package db

import (
	"chathenon/entity"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to the database: %v\n", err)
		return nil
	}

	err = AutoMigrate(db)
	if err != nil {
		fmt.Printf("Error migrating: %v\n", err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		return err
	}

	return nil
}
