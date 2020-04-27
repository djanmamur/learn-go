package main

import (
	"fmt"
	"math"
)

func round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}

func main() {
	var banknote float64
	fmt.Scan(&banknote)

	availableNotes := []int64{100, 50, 20, 10, 5, 2}
	availableCoins := []int64{100, 50, 25, 10, 5, 1}

	notes := int64(banknote)
	coins := int64(round(banknote*float64(100.0))) % 100

	fmt.Println("NOTAS:")
	for _, note := range availableNotes {
		currentNote := notes / note
		notes = notes % note
		fmt.Printf("%d nota(s) de R$ %d.00\n", currentNote, note)
	}

	if notes%10 != 0 {
		coins += notes * 100
	}

	fmt.Println("MOEDAS:")
	for _, coin := range availableCoins {
		currentCoin := coins / coin
		coins %= coin
		fmt.Printf("%d moeda(s) de R$ %.2f\n", currentCoin, float64(coin)/float64(100.0))
	}
}
