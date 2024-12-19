# Chapter 3: Functions

Functions in Go are essential for structuring and organizing code. They enable reusability and modularity in programming. This chapter explores the declaration, syntax, memory handling, and best practices for using functions in Go.

---

## Function Basics

- **Definition**: Functions are used to separate code into logical, reusable blocks.
- **Syntax**:
  ```go
  func add(x, y int) int {
      return x + y
  }
  ```
  
  - Multiple parameters of the same type can be declared together (e.g., `x, y int`).
  - The return type is declared after the parameter list.

---

## Memory and Variables

### Understanding Memory in Go
- Variables in Go store data in memory.
- Assigning a new variable copies the value to a new memory location.

#### Example:
```go
x := 5  // Memory allocates value 5
x = 2    // Memory updates x to 2

y := x  // Allocates new memory for y with value of x (2)
y = 1    // Updates y to 1; x remains 2
```

### Why This Matters
- Variables are **passed by value**, not by reference.
  - A copy of the variable is passed, so the original remains unchanged.

#### Example of Passing by Value:
```go
func main() {
    x := 5
    increment(x)
    fmt.Println(x)  // Output: 5
}

func increment(x int) {
    x++  // This modifies the copy, not the original
}
```

---

## Handling Unused Variables
- **Go does not allow unused variables**: Any variable declared but not used will cause a compilation error.
- **Ignoring return values**: Use the underscore (`_`) to ignore values you don’t need.

#### Example:
```go
result, _ := someFunction() // Ignores the second return value
```

---

## Named Return Values
- **Definition**: A feature where return values are declared in the function signature.
- **Purpose**: Useful for documentation and understanding small, simple functions.

### Example:
```go
func getCoords() (x, y int) {
    x, y = 5, 6
    return  // Implicit return ("naked return")
}
```

### Notes on Named Returns:
- **Advantages**:
  - Improves clarity for small functions.
  - Makes it clear what each returned value represents.
- **Disadvantages**:
  - Can harm readability in complex functions.
  - Explicit returns are generally preferred.

#### Explicit Return Example:
```go
func getCoords() (x, y int) {
    return 5, 6  // Explicit return
}
```

---

## Guard Clauses (Early Returns)
- **Definition**: An early return in a function, typically used to handle errors or edge cases.
- **Purpose**: Provides a linear and readable logic flow.

### Example:
```go
func process(data int) error {
    if data < 0 {
        return fmt.Errorf("Invalid data: %d", data)  // Guard clause
    }
    // Happy path logic here
    return nil
}
```

### Benefits of Guard Clauses:
1. Simplifies complex logic trees.
2. Ensures errors are handled early, allowing the "happy path" to remain clear.

---

## Summary of Best Practices
1. Use named return values for small, simple functions to improve documentation.
2. Prefer explicit returns for clarity in larger functions.
3. Utilize guard clauses to handle errors and edge cases early.
4. Always use variables, or explicitly ignore them, to avoid compilation errors.
5. Remember that Go passes variables by value, so changes in functions do not affect the original variables unless explicitly returned or pointers are used.

