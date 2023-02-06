package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, c float64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	fmt.Sscanf(line, "%f %f %f", &a, &b, &c)

	discriminant := calculateDiscriminant(a, b, c)
	if math.IsNaN(discriminant) || a*2 == 0.0 {
		fmt.Println("Impossivel calcular")
	} else {
		root1, root2 := calculateRoots(a, b, discriminant)
		fmt.Printf("R1 = %.5f\n", root1)
		fmt.Printf("R2 = %.5f\n", root2)
	}
}

func calculateDiscriminant(a, b, c float64) float64 {
	return math.Sqrt(math.Pow(b, 2) - 4*a*c)
}

func calculateRoots(a, b, discriminant float64) (float64, float64) {
	root1 := (-b + discriminant) / (2.0 * a)
	root2 := (-b - discriminant) / (2.0 * a)

	return root1, root2
}
