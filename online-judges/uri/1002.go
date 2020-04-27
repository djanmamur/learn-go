package main

import "fmt"
import "log"

const PI float64 = 3.14159

func main() {
	var radius float64
	_, err := fmt.Scan(&radius)
	if err != nil {
		log.Println("Error reading radius.")
	}

	var area = calculateArea(radius)
	fmt.Printf("A=%.4f\n", area)
}

func calculateArea(radius float64) float64 {
	return PI * (radius * radius)
}