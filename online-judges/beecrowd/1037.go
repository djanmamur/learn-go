package main

import "fmt"

type interval struct {
	low  int64
	high int64
}

func main() {
	var number float64
	fmt.Scan(&number)

	intervalRange := []interval{
		{low: 0, high: 25},
		{low: 25, high: 50},
		{low: 50, high: 75},
		{low: 75, high: 100},
	}

	var currentInterval interval
	for _, intv := range intervalRange {
		if number >= float64(intv.low) && number <= float64(intv.high) {
			currentInterval = intv
			break
		}
	}

	if currentInterval.low == 0 && currentInterval.high == 0 {
		fmt.Println("Fora de intervalo")
	} else {
		openingBracket := "("
		if currentInterval.low == 0 {
			openingBracket = "["
		}
		fmt.Printf("Intervalo %s%d,%d]\n", openingBracket, currentInterval.low, currentInterval.high)
	}
}
