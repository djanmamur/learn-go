package main

import "fmt"

func main() {
	var a, b, c float32

	fmt.Scan(&a, &b, &c)
	average := calculateAverage(a, b, c)
	fmt.Printf("MEDIA = %.1f\n", average)
}

func calculateAverage(a, b, c float32) float32 {
	return (a*2 + b*3 + c*5) / 10.0
}
