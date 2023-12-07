package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// define the timeout of the request
	var timeout time.Duration = 1000 * time.Millisecond

	// create the HTTP client
	client := &http.Client{
		Timeout: timeout,
	}

	// define the URL for the request
	var URL string = "https://jsonplaceholder.typicode.com/posts"

	// create the post
	// this post will be sent to the server
	newPost := struct {
		Title  string
		Body   string
		UserId int
	}{
		Title:  "my new post",
		Body:   "this is the content of the post",
		UserId: 1,
	}

	// parse the post into JSON format
	newPostRequest, err := json.Marshal(newPost)

	// if the parsing process is failed, print out the error
	if err != nil {
		log.Panic(err)
	}

	// create the request body
	requestBody := bytes.NewBuffer(newPostRequest)

	// create the POST request
	req, err := http.NewRequest(http.MethodPost, URL, requestBody)
	if err != nil {
		log.Panic(err)
	}

	// send the request to the server
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	// print out the response body
	fmt.Println("Response body: ", string(body))
	fmt.Println("======")
	// print out the response status code
	fmt.Println("Response Status Code: ", res.StatusCode)
	fmt.Println("======")
	// print out the response headers
	fmt.Println("Response Headers")
	for k, v := range res.Header {
		fmt.Println(k, " : ", v)
	}
}
