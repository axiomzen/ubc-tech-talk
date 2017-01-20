package main

import (
	"log"

	"encoding/json"
	"io"

	"net/http"

	"os"

	"github.com/gocraft/web"
)

var apiToken = os.Getenv("API_TOKEN")

// MyContext will keep track of the context for each request
type MyContext struct{}

func (c *MyContext) render(status int, value interface{}, rw web.ResponseWriter) {
	if value != nil {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(status)
		if err := c.encode(value, rw); err != nil {
			log.Println(err)
		}
	} else {
		rw.WriteHeader(status)
	}
}

func (c *MyContext) encode(value interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(value)
}

func (c *MyContext) decode(value interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(value)
}

// NotFound returns status 404 if the route is not available
func (c *MyContext) NotFound(rw web.ResponseWriter, req *web.Request) {
	c.render(http.StatusNotFound, Message{Message: "Not Found"}, rw)
}

// APIAuthRequired simple api gating
func (c *MyContext) APIAuthRequired(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	token := req.Header.Get("x-api-token")
	if token == apiToken {
		next(rw, req)
	} else {
		c.render(http.StatusUnauthorized, Message{Message: "Unauthorized"}, rw)
	}
}
