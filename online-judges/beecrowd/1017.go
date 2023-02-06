package main

import "fmt"

func main() {
	var timeSpent, averagaSpeed float64
	const carConsumption = 12.0 // 12 km/l
	fmt.Scan(&timeSpent, &averagaSpeed)
	fuelSpent := timeSpent * averagaSpeed / carConsumption
	fmt.Printf("%.3f\n", fuelSpent)
}
