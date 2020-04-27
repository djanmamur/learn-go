package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c, d int64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	fmt.Sscanf(line, "%d %d %d %d", &a, &b, &c, &d)

	if b > c && d > a && c+d > a+b && c > 0 && d > 0 && a%2 == 0 {
		fmt.Println("Valores aceitos")
	} else {
		fmt.Println("Valores nao aceitos")
	}
}
