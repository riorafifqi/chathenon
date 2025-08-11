package db

import (
	"database/sql"
	"fmt"
	"os"

	
	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB() *sql.DB {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("pgx", connStr)

	if err != nil {
		fmt.Printf("Error connecting to the database: %v\n", err)
		return nil
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error pinging the database: %v\n", err)
		return nil
	}

	return db
}
