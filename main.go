package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func init() {
	tmpDB, err := sql.Open("postgres", "user=runhknsftmpuqt dbname=da1qg9a3lcbj6g host=ec2-54-235-240-92.compute-1.amazonaws.com password=89634650246cdb1fbac7e8cbaa00973a708eceda063bb8bc1d108f209c716a44")
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
