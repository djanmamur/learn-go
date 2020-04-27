package main

import "fmt"

func main() {
	var banknote int64

	availableNotes := []int64{100, 50, 20, 10, 5, 2, 1}
	fmt.Scan(&banknote)
	fmt.Println(banknote)
	for _, note := range availableNotes {
		curr := banknote / note
		banknote = banknote % note
		fmt.Printf("%d nota(s) de R$ %d,00\n", curr, note)
	}

}
