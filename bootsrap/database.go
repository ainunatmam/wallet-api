package bootstrap

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DATABASE_USERNAME"),        
		os.Getenv("DATABASE_PASSWORD"),    
		os.Getenv("DATABASE_HOST"),   
		os.Getenv("DATABASE_PORT"),        
		os.Getenv("DATABASE_NAME"),   
	)

	var err error
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB:", err)
	}
	return db, nil
}
