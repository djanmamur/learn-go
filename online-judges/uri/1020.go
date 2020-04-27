package main

const numberOfDaysInYear int32 = 365
const numberOfDaysInMonth int32 = 30

func main() {
	
}

func calculateNumberOfDays(age int32) (int32, int32, int32) {
	years := age / 365
	months := age / 30
	var days int32 = age
	return years, months, days
}
