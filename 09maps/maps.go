package main

import "fmt"

func main() {
	fmt.Println("Go tutorial for Maps")
	mp := make(map[string]string)
	mp["rohan"] = "patel"
	mp["hello"] = "world"
	mp["hotel"] = "trivigo"
	mp["hi"] = "bye"
	fmt.Println("Map is :", mp)
	fmt.Println("last name of rohan is", mp["rohan"])
	delete(mp, "hi")
	fmt.Println(mp)

	for i, v := range mp {
		fmt.Println(i, ":", v)
	}
}
