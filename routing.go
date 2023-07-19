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
	const query = `SELECT id, inside, radiator, outside, added 
				   FROM temperature 
				   ORDER BY id DESC 
				   LIMIT 1`
	row := db.QueryRow(query)
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
