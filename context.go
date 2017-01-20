package main

import (
	"log"

	"encoding/json"
	"io"

	"github.com/gocraft/web"
)

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
