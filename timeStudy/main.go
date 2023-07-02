package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study!")

	currentTime := time.Now()

	fmt.Println(currentTime)

	fmt.Println(currentTime.Format("01 02 2006 Monday 15:04:05 "))

	createdAt := time.Date(2000, time.December, 13, 4, 45, 13, 0, time.UTC)

	fmt.Println(createdAt)

	fmt.Println(createdAt.Format("01 02 2006 Monday"))
}
