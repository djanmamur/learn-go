package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	sides := []int{a, b, c} // Much easier with sorted array of input
	sort.Ints(sides)
	if !isTriangle(sides) {
		fmt.Println("Invalido")
	} else {
		if isEquilateral(sides) {
			fmt.Println("Valido-Equilatero")
		} else if isIsoceles(sides) {
			fmt.Println("Valido-Isoceles")
		} else if isScalene(sides) {
			fmt.Println("Valido-Escaleno")
		}

		if isRectangular(sides) {
			fmt.Println("Retangulo: S")
		} else {
			fmt.Println("Retangulo: N")
		}
	}
}

func isTriangle(sides []int) bool {
	if sides[0]+sides[1] > sides[2] {
		return true
	}
	return false
}

func isEquilateral(sides []int) bool {
	if sides[0] == sides[1] && sides[1] == sides[2] {
		return true
	}
	return false
}

func isScalene(sides []int) bool {
	if sides[0] != sides[1] && sides[1] != sides[2] {
		return true
	}
	return false
}

func isIsoceles(sides []int) bool {
	if sides[0] == sides[1] || sides[1] == sides[2] {
		return true
	}
	return false
}

func isRectangular(sides []int) bool {
	a := float64(sides[0])
	b := float64(sides[1])
	c := float64(sides[2])

	if math.Pow(a, 2)+math.Pow(b, 2) == math.Pow(c, 2) {
		return true
	}
	return false
}
