package main

import "fmt"

func main() {
	var secondsPassed int64
	fmt.Scan(&secondsPassed)

	hour := secondsPassed / 3600
	secondsPassed %= 3600
	minutes := secondsPassed / 60
	secondsPassed %= 60
	seconds := secondsPassed

	fmt.Printf("%d:%d:%d\n", hour, minutes, seconds)

}
