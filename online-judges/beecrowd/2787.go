package main

import "fmt"

func main() {
	var rows, columns int

	fmt.Scan(&rows, &columns)

	// Do we really need to have a loop here? :) Seems to be pretty easy
	sumRowsColums := rows + columns

	if sumRowsColums%2 == 0 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}

}
