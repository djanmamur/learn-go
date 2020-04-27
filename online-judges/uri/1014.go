package main

import "fmt"

func main() {
	var totalDistance, fuelSpent float64

	fmt.Scan(&totalDistance, &fuelSpent)

	// 14.286 km/l
	fmt.Printf("%.3f km/l\n", totalDistance/fuelSpent)
}
