package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var a, b, c int
	var inputLine string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputLine = scanner.Text()
	fmt.Sscanf(inputLine, "%d %d %d", &a, &b, &c)

	nrs := []int{a, b, c}
	sort.Ints(nrs)
	fmt.Printf("%d eh o maior\n", nrs[2])
}
