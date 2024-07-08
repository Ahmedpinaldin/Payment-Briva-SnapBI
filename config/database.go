package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

func NewConnection() (*sql.DB, error) {
	// Replace the connection parameters with your ownerrenv := godotenv.Load()
	errenv := godotenv.Load()

	if errenv != nil {
		// panic("Failed To Load ENV")
		log.Fatalf("Error loading .env file")
	}

	passwordDb := os.Getenv("DB_PASS")
	userDb := os.Getenv("DB_USER")
	portDb := os.Getenv("DB_PORT")
	hostDb := os.Getenv("DB_HOST")
	nameDb := os.Getenv("DB_NAME")

	// Define the connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", hostDb, userDb, passwordDb, portDb, nameDb)

	// Connect to the database
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	return db, nil
}
