package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var productCode1, units1 int
	var pricePerUnit1 float64

	var productCode2, units2 int
	var pricePerUnit2 float64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line1 := scanner.Text()
	scanner.Scan()
	line2 := scanner.Text()
	fmt.Sscanf(line1, "%d %d %f", &productCode1, &units1, &pricePerUnit1)
	fmt.Sscanf(line2, "%d %d %f", &productCode2, &units2, &pricePerUnit2)
	totalPrice := calculatePrice(units1, units2, pricePerUnit1, pricePerUnit2)
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", totalPrice)

}

func calculatePrice(units1, units2 int, pricePerUnit1, pricePerUnit2 float64) float64 {
	return float64(units1)*pricePerUnit1 + float64(units2)*pricePerUnit2
}
