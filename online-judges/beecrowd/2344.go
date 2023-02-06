package main

import "fmt"

func main() {
	var testScore int

	fmt.Scan(&testScore)

	switch {
	case testScore == 0:
		fmt.Println("E")
	case testScore <= 35:
		fmt.Println("D")
	case testScore <= 60:
		fmt.Println("C")
	case testScore <= 85:
		fmt.Println("B")
	default:
		fmt.Println("A")
	}
}
