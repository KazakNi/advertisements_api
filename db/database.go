package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var DB = InitDB()

func InitDB() *sqlx.DB {
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
	return db
}

func CreateDB(db *sqlx.DB) {
	db.MustExec(schema)
}

func ExecuteQueries(db *sqlx.DB) {
	// For database filling, disabled while prod mode
	tx := db.MustBegin()
	log.Println("Query transaction begins")
	tx.MustExec(insert_query, "Balls", "Adiddas premium match ball", "150", Creation_time, `{"link11", "link22", "link34"}`)
	tx.MustExec(insert_query, "Glasses", "Oakley used but good", "70", Creation_time, `{"link1", "link2", "link3"}`)
	tx.MustExec(insert_query, "Tutor", "English tutor for children, price per hour", "50", Creation_time, `{"http://image.com/asfalm122.jpeg", "http://image.com/asfaf345hjtuj67.jpeg", "http://image.com/vbSDGQ93m.jpeg"}`)
	log.Println("Query transaction ends")
	tx.Commit()
}

func CountRows(db *sqlx.DB) int {
	rows, err := db.Query("SELECT COUNT(*) FROM advertisements")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
			fmt.Println(count)
		}
	}
	return count
}
