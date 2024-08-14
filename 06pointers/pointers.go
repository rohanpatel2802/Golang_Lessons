package main

import "fmt"

func main() {
	fmt.Println("Working with pointers")
	var ptr *int
	fmt.Println("Default value of interger pointer is ", ptr)

	myNumber := 28
	ptr = &myNumber
	fmt.Println("Value of ptr :", ptr)
	fmt.Println("Value at ptr :", *ptr)
}
