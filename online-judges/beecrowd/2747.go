package main

import "fmt"

func main() {
	printCharacter("-", 38)
	for i := 0; i <= 4; i++ {
		printLine(38)
	}
	printCharacter("-", 38)
}

func printCharacter(character string, numDashes int) {
	for i := 0; i <= numDashes; i++ {
		fmt.Print(character)
	}
	fmt.Println()
}

func printLine(numLine int) {
	for i := 0; i <= numLine; i++ {
		if i == 0 || i == numLine {
			fmt.Print("|")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
