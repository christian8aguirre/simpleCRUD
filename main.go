package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func init() {
	tmpDB, err := sql.Open("postgres", "url")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	fs := http.FileServer(http.Dir("/www/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", handleListBooks)
	http.HandleFunc("/book.html", handleViewBook)
	http.HandleFunc("/save", handleSaveBook)
	http.HandleFunc("/delete", handleDeleteBook)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
