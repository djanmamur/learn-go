package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)
	floorWorkers := []int{a, b, c}

	var smallest = 100000000
	for i := 1; i <= 3; i++ {
		stepsForFloor := countStepsForFloor(i, floorWorkers)
		if stepsForFloor < smallest {
			smallest = stepsForFloor
		}
	}

	fmt.Println(smallest)
}

func countStepsForFloor(currentFloor int, floorWorkers []int) int {
	const STEPS_PERSON = 2.0
	var steps = 0
	for i := 1; i <= 3; i++ {
		if currentFloor == i {
			continue
		}

		steps += int(STEPS_PERSON * math.Abs(float64(currentFloor)-float64(i)) * float64(floorWorkers[i-1]))
	}

	return steps
}
