package main

import (
	"net/http"

	"strconv"

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

// GetQuestions gets a list of questions
func (c *MyContext) GetQuestions(rw web.ResponseWriter, req *web.Request) {
	var questions Questions
	if err := db.Model(&questions).Select(); err != nil {
		c.render(http.StatusBadRequest, nil, rw)
		return
	}
	c.render(http.StatusOK, questions, rw)
}

// AskQuestions saves a question to the db
func (c *MyContext) AskQuestions(rw web.ResponseWriter, req *web.Request) {
	var question Question
	if err := c.decode(&question, req.Body); err != nil {
		c.render(http.StatusBadRequest, nil, rw)
		return
	}
	if err := db.Insert(&question); err != nil {
		c.render(http.StatusBadRequest, nil, rw)
		return
	}
	c.render(http.StatusOK, question, rw)
}

// Upvote allows adding an upvote to the question
func (c *MyContext) Upvote(rw web.ResponseWriter, req *web.Request) {
	questionID, _ := strconv.ParseInt(req.PathParams["question_id"], 10, 64)
	question := Question{
		ID: questionID,
	}
	if res, err := db.Model(&question).Set("votes = votes + 1").Where("id = ?id").Update(); err != nil {
		c.render(http.StatusBadRequest, Message{Message: "Malformed body"}, rw)
		return
	} else if res.RowsAffected() != 1 {
		c.render(http.StatusNotFound, Message{Message: "Question not found"}, rw)
		return
	}
	c.render(http.StatusNoContent, nil, rw)
}
