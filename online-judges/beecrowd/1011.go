package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Scan(&radius)
	sphereVolume := calculateSphereVolume(radius)
	fmt.Printf("VOLUME = %.3f\n", sphereVolume)

}

func calculateSphereVolume(radius float64) float64 {
	const PI = 3.14159
	return 4.0 / 3 * PI * math.Pow(radius, 3)
}
