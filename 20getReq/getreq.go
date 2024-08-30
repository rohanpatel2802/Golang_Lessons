package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Go tutorial on GET request")
	makeGetRequest()

}

func makeGetRequest() {
	const url = "http://localhost:8000/get"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Status code:", resp.Status)
}
