package main

import "fmt"

func main() {
	for i := 97; i < 123; i++ {
		fmt.Printf("%d %s %s\n", i, "e", toCharStr(i))
	}
}

func toCharStr(i int) string {
	return string('a' - 97 + i)
}
