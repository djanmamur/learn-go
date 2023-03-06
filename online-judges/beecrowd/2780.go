package main

import "fmt"

func main() {
	var distance int

	fmt.Scan(&distance)

	switch {
	case distance <= 800:
		fmt.Println(1)
	case distance <= 1400:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}
}
