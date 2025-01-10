package models

import (
	"database/sql"
	"fmt" // 確保引入 fmt 包

	_ "github.com/lib/pq"
)

func GetDBConnection() (*sql.DB, error) {
    connStr := "host=postgres-data user=postgres password=Sammy0928 dbname=violations_db sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("Failed to open the database: %v", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("Failed to ping the database: %v", err)
    }

    return db, nil
}
