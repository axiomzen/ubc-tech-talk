package main

// Question will represent our questions
type Question struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	Author string `json:"author"`
	Votes  int32  `json:"votes"`
}

// Questions a slice of Question pointers
type Questions []*Question

// Message will represent the messages we will be sending back in some responses
type Message struct {
	Message string `json:"message"`
}
