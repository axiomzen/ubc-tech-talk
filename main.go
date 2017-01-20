package main

import (
	"log"
	"net/http"
	"os"

	"net/url"

	"path"

	"github.com/gocraft/web"
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

	router := web.New(MyContext{})
	router.Middleware(web.LoggerMiddleware)
	router.Middleware(web.ShowErrorsMiddleware)
	staticPath, _ := os.Getwd()
	if os.Getenv("ENVIRONMENT") == "production" {
		staticPath = "/app"
	}
	router.Middleware(web.StaticMiddleware(path.Join(staticPath, "public"), web.StaticOption{IndexFile: "index.html"}))
	router.Get("/ping", (*MyContext).Ping)

	subRouter := router.Subrouter(MyContext{}, "")
	subRouter.Middleware((*MyContext).APIAuthRequired)

	subRouter.Get("/questions", (*MyContext).GetQuestions)
	subRouter.Post("/questions", (*MyContext).AskQuestions)
	subRouter.Post("/questions/:question_id:\\d*/upvote", (*MyContext).Upvote)

	port := os.Getenv("PORT")

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}
