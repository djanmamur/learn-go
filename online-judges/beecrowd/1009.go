package main

import "fmt"

func main() {
	var employeeName string
	var monthlySalary, totalSales float64
	fmt.Scan(&employeeName, &monthlySalary, &totalSales)

	employeeSalary := calculateSalary(monthlySalary, totalSales)
	fmt.Printf("TOTAL = R$ %.2f\n", employeeSalary)

}

func calculateSalary(monthlySalary, totalSales float64) float64 {
	// Seller receives 15% over all products sold
	const SalePayPercentage = .15

	return monthlySalary + totalSales*SalePayPercentage
}
