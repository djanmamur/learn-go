package main

import (
	"fmt"
	"strings"
)

func main() {
	// Shitty solution here, no offence to myself
	// Could be calculated in a function; come back to this later definitely

	fmt.Printf("%s\n", strings.Repeat("-", 39))
	fmt.Printf("|%15s%22s|\n", "Roberto", " ")
	fmt.Printf("|%s|\n", strings.Repeat(" ", 37))
	fmt.Printf("|%12s%25s|\n", "5786", " ")
	fmt.Printf("|%s|\n", strings.Repeat(" ", 37))
	fmt.Printf("|%14s%23s|\n", "UNIFEI", " ")
	fmt.Printf("%s\n", strings.Repeat("-", 39))
}
