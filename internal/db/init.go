package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Db struct {
	db *sqlx.DB
}

func NewDataBase() *Db {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("user=%s password=%s host=%s dbname=%s",
			dbUser,
			dbPassword,
			dbHost,
			dbName,
		),
	)
	if err != nil {
		log.Panicf("could not connect to PostgreSQL. error: %v", err)
	}

	return &Db{db: db}
}
