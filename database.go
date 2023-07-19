package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"os"
	"time"
)

var db *sql.DB

type Temperature struct {
	Id       int       `json:"-"`
	Inside   float32   `json:"inside"`
	Radiator float32   `json:"radiator"`
	Outside  float32   `json:"outside"`
	Added    time.Time `json:"added"`
}

func initDB() io.Closer {
	var err error
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	db, err = sql.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			dbUser, dbPass, dbHost, dbPort, dbName,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
