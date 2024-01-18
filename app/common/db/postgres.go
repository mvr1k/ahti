package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewDbInstance(host string, port int, user string, password string, database string) *sql.DB {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database))
	if err != nil {
		panic(err)
	}

	return db

}
