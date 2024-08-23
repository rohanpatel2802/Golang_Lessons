package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?courseName=js&paymentId=12345"

func main() {
	fmt.Println("Go tutorial for urls")
	fmt.Println("My url is", myurl)
	res, _ := url.Parse(myurl)
	fmt.Printf("Result is of type %T\n", res)
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	qparams := res.Query()
	fmt.Printf("Type of qparams is : %T\n", qparams)
	fmt.Println(qparams)

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "rohan.dev:1234",
		Path:    "cdot",
		RawPath: "user=rp",
	}
	anotherUrl := partsofUrl.String()
	fmt.Println(anotherUrl)

}
