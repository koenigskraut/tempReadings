package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("https://%s:443%s", domain, r.RequestURI), http.StatusMovedPermanently)
}

func getRoot(writer http.ResponseWriter, _ *http.Request) {
	b, _ := content.ReadFile("static/mainPage.html")
	_, _ = writer.Write(b)
}

func getLastReading(w http.ResponseWriter, _ *http.Request) {
	row := db.QueryRow(LastReadingQuery)
	if row == nil {
		w.Write([]byte("{}"))
		return
	}
	var t Temperature
	if err := row.Scan(&t.Id, &t.Inside, &t.Radiator, &t.Outside, &t.Added); err != nil {
		log.Println(err)
	}
	if err := json.NewEncoder(w).Encode(&t); err != nil {
		log.Println(err)
	}
}

type LimitOffset struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ReadingsType int

const (
	LastReadings ReadingsType = iota
	FirstReadings
)

func getNReadings(rt ReadingsType) func(w http.ResponseWriter, r *http.Request) {
	var query string
	if rt == LastReadings {
		query = LastNReadingsQuery
	} else {
		query = FirstNReadingsQuery
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var args LimitOffset
		err := json.NewDecoder(r.Body).Decode(&args)
		if err != nil {
			log.Println("Reading data error:", err)
			w.Write([]byte("[]"))
			return
		}
		scanned := make([]Temperature, 0, args.Limit)
		rows, err := db.Query(query, args.Limit, args.Offset)
		if err != nil {
			log.Println("Query error:", err)
		}
		if rows == nil {
			w.Write([]byte("[]"))
			return
		}
		defer rows.Close()
		for rows.Next() {
			var t Temperature
			if err := rows.Scan(&t.Id, &t.Inside, &t.Radiator, &t.Outside, &t.Added); err != nil {
				log.Println(err)
			}
			scanned = append(scanned, t)
		}
		if err := json.NewEncoder(w).Encode(&scanned); err != nil {
			log.Println(err)
		}
	}
}
