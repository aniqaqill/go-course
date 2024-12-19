package main

import "fmt"

// func main() {
// chaptertwo()
// shortvariable()
// conditionals()
// }

func chaptertwo() {
	// Boolean
	var isGoAwesome bool = true
	fmt.Println(isGoAwesome)

	// String
	var greeting string = "Hello, Go!"
	fmt.Println(greeting)

	// Integer
	var age int = 30
	var smallNumber int8 = 127
	fmt.Println(age, smallNumber)

	// Unsigned Integer
	var positiveNumber uint = 42
	var byteValue byte = 255
	fmt.Println(positiveNumber, byteValue)

	// Rune
	var unicodeChar rune = 'G'
	fmt.Println(unicodeChar)

	// Floating Point
	var pi float32 = 3.14
	var e float64 = 2.718281828459045
	fmt.Println(pi, e)

	// Complex Numbers
	var complexNumber complex64 = 1 + 2i
	var complexNumber2 complex128 = 2 + 3i
	fmt.Println(complexNumber, complexNumber2)
}

//short assignment operator

func shortvariable() {
	// Boolean
	isGoAwesome := true
	fmt.Println(isGoAwesome)

	// String
	greeting := "Hello, Go!"
	fmt.Println(greeting)

	// Integer
	age := 30
	smallNumber := int8(127)
	fmt.Println(age, smallNumber)

	// Unsigned Integer
	positiveNumber := uint(42)
	byteValue := byte(255)
	fmt.Println(positiveNumber, byteValue)

	// Rune
	unicodeChar := 'G'
	fmt.Println(unicodeChar)

	// Floating Point
	pi := float32(3.14)
	e := 2.718281828459045
	fmt.Println(pi, e)

	// Complex Numbers
	complexNumber := complex(1, 2)
	complexNumber2 := complex(2, 3)
	fmt.Println(complexNumber, complexNumber2)
}

func conditionals() {
	messagelen := 20
	if messagelen > 10 {
		fmt.Println("Message is too long")
	} else {
		fmt.Println("Message length is acceptable")
	}
}
