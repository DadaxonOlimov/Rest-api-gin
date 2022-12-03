package connaction

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Connection() *sql.DB {

	connStr := "user=postgres password=1217 dbname=crud sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
