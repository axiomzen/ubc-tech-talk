package main

import (
	"net/http"

	"github.com/gocraft/web"
)

// Ping checks the status of the server
func (c *MyContext) Ping(rw web.ResponseWriter, req *web.Request) {
	if _, err := db.ExecOne("SELECT 1"); err != nil {
		c.render(http.StatusBadRequest, nil, rw)
		return
	}
	c.render(http.StatusOK, Message{Message: "pong"}, rw)
}
