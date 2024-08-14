package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current time : ", currentTime)
	formattedTime := currentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println("Current time in new format: ", formattedTime)
	createdDate := time.Date(2000, time.January, 28, 7, 28, 0, 0, time.UTC)
	fmt.Println("Created Date and time = ", createdDate.Format("01-02-2006 15:04:05 Monday"))
	time.Sleep(time.Duration(time.Minute))
}
