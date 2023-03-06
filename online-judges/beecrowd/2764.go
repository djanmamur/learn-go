package main

import (
	"fmt"
	"strings"
)

func main() {
	var date string
	fmt.Scan(&date)

	dateParts := strings.Split(date, "/")
	day := dateParts[0]
	month := dateParts[1]
	year := dateParts[2]

	fmt.Printf("%s/%s/%s\n", month, day, year)
	fmt.Printf("%s/%s/%s\n", year, month, day)
	fmt.Printf("%s-%s-%s\n", day, month, year)
}
