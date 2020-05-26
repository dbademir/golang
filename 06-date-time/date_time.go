package main

import (
	"fmt"
	"time"
)

func main() {

	userTime := time.Date(2020, time.May, 26, 23, 0, 0, 0, time.UTC)

	fmt.Printf("User defined datetime%s\n\n", userTime)

	now := time.Now()
	fmt.Printf("Current datetime is %s\n\n", now)

	fmt.Println("User defined month is", userTime.Month())
	fmt.Println("User defined day is", userTime.Day())
	fmt.Println("User defined weekday is", userTime.Weekday())

	tomorrow := userTime.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n\n",
		tomorrow.Weekday(),
		tomorrow.Month(),
		tomorrow.Day(),
		tomorrow.Year())

	longDateFormat := "Wednesday, May 27, 2020"
	fmt.Println("Tomorrow long format is", tomorrow.Format(longDateFormat))

	fmt.Println("")

	shortFormat := "05/26/20"
	fmt.Println("Tomorrow short format is", tomorrow.Format(shortFormat))
}
