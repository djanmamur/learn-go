package main

import "fmt"

func main() {
	var employeeNumber, hoursWorked int
	var hourlyRate float32

	fmt.Scan(&employeeNumber, &hoursWorked, &hourlyRate)
	employeeSalary := calculateSalary(hoursWorked, hourlyRate)
	/*
		Output example

		NUMBER = 25
		SALARY = U$ 550.00
	*/
	fmt.Printf("NUMBER = %d\n", employeeNumber)
	fmt.Printf("SALARY = U$ %.2f\n", employeeSalary)

}

func calculateSalary(hoursWorked int, hourlyRate float32) float32 {
	return float32(hoursWorked) * hourlyRate
}
