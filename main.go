package main

import (
	"embed"
	"log"
	"net/http"
	"os"
)

//go:embed static/*
var content embed.FS

var domain, certFile, keyFile string

func main() {
	domain = os.Getenv("DOMAIN")
	certFile = os.Getenv("CERT_FILE")
	keyFile = os.Getenv("KEY_FILE")

	toClose := initDB()
	defer toClose.Close()

	go runUDPServer()

	fs := http.FileServer(http.FS(content))
	mux := http.NewServeMux()
	mux.Handle("/static/", fs)
	mux.HandleFunc("/api/lastReading/", getLastReading)
	mux.HandleFunc("/", getRoot)

	go func() {
		if err := http.ListenAndServe(":80", http.HandlerFunc(redirectTLS)); err != nil {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()
	err := http.ListenAndServeTLS(":443", certFile, keyFile, mux)
	if err != nil {
		log.Fatal(err)
	}
}
