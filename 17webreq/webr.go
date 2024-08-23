package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.example.com"

func main() {
	fmt.Println("Tutorial for web request")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Response is of type %T\n", res)
		fmt.Println("Response is : ", res)
		defer res.Body.Close()
		databytes, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			panic(err2)
		} else {
			content := string(databytes)
			fmt.Println(content)
		}
	}

}
