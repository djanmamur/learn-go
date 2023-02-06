package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var a, b, c, d float64

	fmt.Sscanf(scanner.Text(), "%f %f %f %f", &a, &b, &c, &d)
	media := (a*2 + b*3 + c*4 + d) / 10
	fmt.Printf("Media: %.1f\n", media)

	if media >= 7.0 {
		fmt.Println("Aluno aprovado.")
	} else if media >= 5.0 {
		fmt.Println("Aluno em exame.")
		var additionalGrade float64
		fmt.Scan(&additionalGrade)
		fmt.Printf("Nota do exame: %.1f\n", additionalGrade)

		newMedia := (media + additionalGrade) / 2.0
		if newMedia >= 5.0 {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado.")
		}

		fmt.Printf("Media final: %.1f\n", newMedia)
	} else {
		fmt.Println("Aluno reprovado.")
	}
}
