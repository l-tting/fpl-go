package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("pgx", "postgres://postgres:6979@localhost:5432/fpl_db")

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil{
		log.Fatal("Cannot connect to db",err)
	}

	return db
}
