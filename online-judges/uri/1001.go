package main

import (
	"fmt"
	"log"
)

func main() {
	var x int
	var y int

	_, err := fmt.Scan(&x, &y)
	if err != nil {
		log.Print("Error reading x and y.")
	}

	result := solve(x, y)
	fmt.Printf("X = %d\n", result)
}

func solve(x int, y int) int {
	return x + y
}
