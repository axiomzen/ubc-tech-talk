package main

import (
	"log"
	"net/http"
	"os"

	"net/url"

	pg "gopkg.in/pg.v5"
)

var db *pg.DB

func main() {

	dbURL, _ := url.Parse(os.Getenv("DATABASE_URL"))

	dbOpts := &pg.Options{
		User:     dbURL.User.Username(),
		Database: dbURL.Path[1:],
		Addr:     dbURL.Host,
	}

	if dbPass, ok := dbURL.User.Password(); ok {
		dbOpts.Password = dbPass
	}

	db = pg.Connect(dbOpts)

	port := os.Getenv("PORT")

	server := http.Server{
		Addr: ":" + port,
	}
	log.Fatal(server.ListenAndServe())
}
