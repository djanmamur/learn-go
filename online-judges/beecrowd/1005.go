package main

import (
	"fmt"
	"log"
)

func main() {
	var a float32
	var b float32

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		log.Println("Error reading `a` and `b`")
	}

	average := calculateAverage(a, b)
	fmt.Printf("MEDIA = %.5f\n", average)
}

func calculateAverage(a float32, b float32) float32 {
	var weightA float32 = 3.5
	var weightB float32 = 7.5
	return (a*weightA + b*weightB) / (weightA + weightB)
}
