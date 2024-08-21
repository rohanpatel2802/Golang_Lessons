package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("tutorial for loops")

	day := make([]string, 0)
	day = append(day, "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday")
	fmt.Println(day)

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Current generated day is", day[rand.Intn(7)])
	// for d := 0; d < len(day); d++ {
	// 	fmt.Println(day[d])
	// }

	// for i := range day {
	// fmt.Println(day[i])
	// }

	// for index, value := range day {
	// 	fmt.Println("Index is", index, " and value is", value)
	// }

}
