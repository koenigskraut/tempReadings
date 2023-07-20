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

// InsertQuery inserts readings into the table, accepts three float32 arguments
const InsertQuery = `
INSERT INTO temperature (inside, radiator, outside) VALUES (?,?,?)
`

// LastReadingQuery fetches last temperature reading, no arguments required
const LastReadingQuery = `
SELECT id, inside, radiator, outside, added 
FROM temperature 
ORDER BY id DESC 
LIMIT 1
`

// LastNReadingsQuery fetches last N temperature readings with offset O, N and O are integer arguments
const LastNReadingsQuery = `
SELECT * FROM (
	SELECT id, inside, radiator, outside, added
	FROM temperature
	ORDER BY id DESC
	LIMIT ? OFFSET ?
) _
ORDER BY id ASC
`

// FirstNReadingsQuery fetches first N temperature readings with offset O, N and O are integer arguments
const FirstNReadingsQuery = `
SELECT id, inside, radiator, outside, added
FROM temperature
ORDER BY id ASC
LIMIT ? OFFSET ?
`

// AverageByTimeQuery fetches all data averaged by specified time interval T in seconds,
// T is an integer argument (accepted twice)
const AverageByTimeQuery = `
SELECT AVG(t.inside) inside, AVG(t.radiator) radiator, AVG(t.outside) outside, FROM_UNIXTIME(t.added) added
FROM (
	SELECT inside, radiator, outside, FLOOR(UNIX_TIMESTAMP(added)/?) * ? added
	FROM temperature
	ORDER BY added ASC
) t
GROUP BY t.added
LIMIT 2000
`

// MinMaxByTimeQuery fetches all data min-maxed by specified time interval T in seconds,
// that is, for given interval T it returns min and max values on it,
// T is an integer argument (accepted twice)
const MinMaxByTimeQuery = `
SELECT 
    MIN(t.inside) inside_min, MAX(t.inside) inside_max, 
    MIN(t.radiator) radiator_min, MAX(t.radiator) radiator_max,
    MIN(t.outside) outside_min, MAX(t.outside) outside_max,
    FROM_UNIXTIME(t.added) added
FROM (
	SELECT inside, radiator, outside, FLOOR(UNIX_TIMESTAMP(added)/?) * ? added
	FROM temperature
	ORDER BY added ASC
) t
GROUP BY t.added
LIMIT 2000
`

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
