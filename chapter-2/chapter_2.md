# Chapter 2: Variables

Go's basic variable types and their usage are key to building robust applications. This guide provides an overview of variable types, usage examples, and additional tips to master variable handling in Go.

---

## Boolean
- **bool**: Represents a boolean value, which can be either `true` or `false`.
  
### Example:
```go
var isReady bool = true
fmt.Printf("Is ready: %t\n", isReady) // Output: Is ready: true
```

---

## String
- **string**: Represents a sequence of characters. Strings are immutable in Go.
  
### Example:
```go
var message string = "Hello, Go!"
fmt.Println(message) // Output: Hello, Go!
```

---

## Integer
- **int**: A signed integer type that is at least 32 bits in size (platform-dependent).
- **int8**: An 8-bit signed integer ranging from -128 to 127.
- **int16**: A 16-bit signed integer ranging from -32768 to 32767.
- **int32**: A 32-bit signed integer ranging from -2147483648 to 2147483647.
- **int64**: A 64-bit signed integer ranging from -9223372036854775808 to 9223372036854775807.
  
### Example:
```go
var age int = 30
fmt.Printf("Age: %d\n", age) // Output: Age: 30
```

---

## Unsigned Integer
- **uint**: An unsigned integer type that is at least 32 bits in size (platform-dependent).
- **uint8**: An 8-bit unsigned integer ranging from 0 to 255.
- **uint16**: A 16-bit unsigned integer ranging from 0 to 65535.
- **uint32**: A 32-bit unsigned integer ranging from 0 to 4294967295.
- **uint64**: A 64-bit unsigned integer ranging from 0 to 18446744073709551615.
- **byte**: Alias for `uint8`, typically used to represent binary data.
  
### Example:
```go
var data byte = 255
fmt.Printf("Data: %d\n", data) // Output: Data: 255
```

---

## Rune
- **rune**: Alias for `int32`, represents a Unicode code point.
  
### Example:
```go
var char rune = 'A'
fmt.Printf("Character: %c, Unicode: %U\n", char, char) // Output: Character: A, Unicode: U+0041
```

---

## Floating Point
- **float32**: A 32-bit IEEE 754 floating-point number.
- **float64**: A 64-bit IEEE 754 floating-point number.
  
### Example:
```go
var pi float64 = 3.14159
fmt.Printf("Pi: %.2f\n", pi) // Output: Pi: 3.14
```

---

## Complex Numbers
- **complex64**: A complex number with `float32` real and imaginary parts.
- **complex128**: A complex number with `float64` real and imaginary parts.
  
### Example:
```go
var z complex128 = complex(1, 2)
fmt.Printf("Complex Number: %v\n", z) // Output: Complex Number: (1+2i)
```

---

## Short Variable Declaration
In Go, you can use the short variable declaration syntax (:=) to declare and initialize variables in one step. This is particularly useful for local variables.
  
### Example:
```go
name := "John"
age := 25
fmt.Printf("Name: %s, Age: %d\n", name, age) // Output: Name: John, Age: 25
```

---

## Checking Variable Types
You can check the type of a variable using the `%T` format specifier in `fmt.Printf`.
  
### Example:
```go
fmt.Printf("Type of age: %T\n", age) // Output: Type of age: int
```

---

## Multiple Variable Declarations
You can declare multiple variables in the same line.
  
### Example:
```go
var x, y, z int = 1, 2, 3
fmt.Printf("x: %d, y: %d, z: %d\n", x, y, z) // Output: x: 1, y: 2, z: 3
```

---

## Type Conversion
Go is strict about type conversions. You need to explicitly convert types when necessary.
  
### Example:
```go
var a int = 42
var b float64 = float64(a)
fmt.Printf("Converted: %f\n", b) // Output: Converted: 42.000000
```

---

## Constants
Constants are immutable values that are known at compile time and do not change during the execution of the program.
  
### Example:
```go
const Pi = 3.14
fmt.Printf("Value of Pi: %f\n", Pi) // Output: Value of Pi: 3.140000
```

---

## String Formatting
Go provides several ways to format strings using the `fmt` package.

### Format Specifiers:
- **%v**: Default format.
- **%T**: Type of the variable.
- **%s**: String.
- **%d**: Decimal integer.
  
### Examples:
#### `fmt.Printf`
Prints formatted string to the console.
```go
fmt.Printf("Hello, %s!\n", "world") // Output: Hello, world!
```

#### `fmt.Sprintf`
Returns a formatted string.
```go
formattedString := fmt.Sprintf("Age: %d", age)
fmt.Println(formattedString) // Output: Age: 25
```

---

## Conditions
### If Statements
In Go, `if` statements do not require parentheses around the condition.
  
### Example:
```go
if height > 4 {
    fmt.Println("You are tall enough!")
}
```

### If with Initial Statement
You can include an initial statement in an `if` block.
  
### Example:
```go
if height := 5; height > 4 {
    fmt.Println("You are tall enough!")
}
```

