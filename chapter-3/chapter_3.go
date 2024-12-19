package main

import "fmt"

func main() {
	result := sub(10, 5)
	fmt.Println(result)
	fmt.Println(concat("Hello", "Go!"))
	x := 5
	fmt.Println(increment(x))
	test(20)
}

func sub(x int, y int) int {
	return x - y
}

func concat(x string, y string) string {
	return x + " " + y
}

func increment(x int) int {
	x++
	return x
}

func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int, yearsUntilRetirement int, yearsUntilSeniorCitizen int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}

	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	yearsUntilRetirement = 65 - age
	if yearsUntilRetirement < 0 {
		yearsUntilRetirement = 0
	}

	yearsUntilSeniorCitizen = 60 - age
	if yearsUntilSeniorCitizen < 0 {
		yearsUntilSeniorCitizen = 0
	}

	return
}

func test(age int) {
	fmt.Println("Age", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental, yearsUntilRetirement, yearsUntilSeniorCitizen := yearsUntilEvents(age)
	fmt.Println("Years until adult:", yearsUntilAdult)
	fmt.Println("Years until drinking:", yearsUntilDrinking)
	fmt.Println("Years until car rental:", yearsUntilCarRental)
	fmt.Println("Years until retirement:", yearsUntilRetirement)
	fmt.Println("Years until senior citizen:", yearsUntilSeniorCitizen)
}
