package main

import "fmt"

func main() {
	fmt.Println("This is an example of defer")
	mydefer()
	defer fmt.Print(" Rohan")
	defer fmt.Print(" is")
	defer fmt.Print(" name")
	fmt.Print("My")

}
func mydefer() {
	fmt.Println("This will print 5 to 1")
	for i := 1; i <= 5; i++ {
		defer fmt.Print(i, " ")
	}
}
