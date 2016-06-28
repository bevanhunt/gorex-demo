package main

import (
	"fmt"
	"time"

	"github.com/bevanhunt/gorex"
)

// JSONReceive - json response
type JSONReceive struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int64  `json:"userId"`
}

// JSONSend - json post
type JSONSend struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int64  `json:"userId"`
}

func main() {
	timeout := 10 * time.Second
	jsonData := &JSONSend{
		Title:  "fancy book",
		Body:   "this is a fancy book",
		UserID: 12,
	}
	req, err := gorex.Request{
		URI:     "http://jsonplaceholder.typicode.com/posts",
		Timeout: timeout}.PostJSON(jsonData)
	if err != nil {
		fmt.Println(err)
	}
	res, err := req.Do()
	if err != nil {
		fmt.Println(err)
	}
	resp := &JSONReceive{}
	res.JSON(resp)
	// should print - this is a fancy book
	fmt.Println(resp.Body)
}
