package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS advertisements(
	
)`

func InitDB() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DriverName := os.Getenv("DBNAME")
	DataSourceName := os.Getenv("DBSOURCE")

	db, err := sqlx.Connect(DriverName, DataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
