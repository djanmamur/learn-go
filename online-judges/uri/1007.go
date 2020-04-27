package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	difference := calculateDifference(a, b, c, d)
	fmt.Printf("DIFERENCA = %d\n", difference)
}

func calculateDifference(a, b, c, d int) int {
	return a*b - c*d
}
