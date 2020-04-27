package main

import (
	"bufio"
	"fmt"
	"os"
)

type product struct {
	code  int64
	price float64
}

func main() {
	var productCode, productQuantity int64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Sscanf(scanner.Text(), "%d %d", &productCode, &productQuantity)

	productMap := map[int64]float64{
		1: 4.00,
		2: 4.50,
		3: 5.00,
		4: 2.00,
		5: 1.50,
	}

	productPrice := productMap[productCode] * float64(productQuantity)
	fmt.Printf("Total: R$ %.2f\n", productPrice)
}
