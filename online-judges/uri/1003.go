package main

import (
	"fmt"
	"log"
)

func main() {
	var a int
	var b int

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		log.Println("Error reading `a` and `b`.")
	}
	product := calculateProduct(a, b)
	fmt.Printf("PROD = %d\n", sum)
}

func calculateProduct(x int, y int) int {
	return x * y
}
