package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Getting input for the user")
	welcome := "Welcome to user input tutorial"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name")
	input, _ := reader.ReadString('\n')
	fmt.Println("So you name is : ", input, " y/n")
	input, _ = reader.ReadString('\n')
}
