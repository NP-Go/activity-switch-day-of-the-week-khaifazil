package main

import (
	"fmt"
	"time"
)

func main() {
	//Insert code here
	today := time.Now().Weekday()

	switch today {
	case time.Monday, time.Wednesday, time.Friday, time.Sunday:
		fmt.Println("Today is", today, "\nIt's an odd day")
	case time.Tuesday, time.Thursday, time.Saturday:
		fmt.Println("Today is", today, "\nIt's an even day")
	default:
		fmt.Println("what day is it?!")
	}
}
