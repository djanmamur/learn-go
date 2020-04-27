package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const PI = 3.14159

type triangle struct {
	base   float64
	height float64
}

type circle struct {
	radius float64
}

type trapezium struct {
	highBase float64
	lowBase  float64
	height   float64
}

type square struct {
	side float64
}

type rectangle struct {
	width  float64
	height float64
}

type shape interface {
	area() float64
}

func (t triangle) area() float64 {
	return t.height * t.base / 2.0
}

func (c circle) area() float64 {
	return PI * math.Pow(c.radius, 2)
}

func (t trapezium) area() float64 {
	return (t.highBase + t.lowBase) / 2.0 * t.height
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	var a, b, c float64
	var inputLine string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputLine = scanner.Text()
	fmt.Sscanf(inputLine, "%f %f %f", &a, &b, &c)

	listShapes := []struct {
		name  string
		shape shape
	}{
		{name: "TRIANGULO", shape: triangle{a, c}},
		{name: "CIRCULO", shape: circle{c}},
		{name: "TRAPEZIO", shape: trapezium{a, b, c}},
		{name: "QUADRADO", shape: square{b}},
		{name: "RETANGULO", shape: rectangle{a, b}},
	}
	for _, shapeStruct := range listShapes {
		fmt.Printf("%s: %.3f\n", shapeStruct.name, shapeStruct.shape.area())
	}
}
