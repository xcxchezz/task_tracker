package db

import (
	"database/sql"
)

func connectDB() (*sql.DB, error) {
	connStr := "user=postgres password=llletMedie5 dbname=db_tracker sslmode=disable"
	return sql.Open("postgres", connStr)
}
