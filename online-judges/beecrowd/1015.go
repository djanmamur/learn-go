package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var x1, x2, y1, y2 float64
	var inputPoint1, inputPoint2 string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputPoint1 = scanner.Text()
	fmt.Sscanf(inputPoint1, "%f %f", &x1, &y1)

	scanner.Scan()
	inputPoint2 = scanner.Text()
	fmt.Sscanf(inputPoint2, "%f %f", &x2, &y2)

	distance := calculateDistance(x1, x2, y1, y2)
	fmt.Printf("%.4f\n", distance)
}

func calculateDistance(x1, x2, y1, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
