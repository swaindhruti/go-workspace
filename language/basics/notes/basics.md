# Go Basics Notes

## Index

- [Overview](#overview)
- [Tooling and Execution](#tooling-and-execution)
- [Packages, Imports, and Standard Library](#packages-imports-and-standard-library)
- [Types, Values, and Naming](#types-values-and-naming)
- [Arrays, Slices, Maps, and Runes](#arrays-slices-maps-and-runes)
- [Functions and Variadic Arguments](#functions-and-variadic-arguments)
- [Defer, Panic, and Recover](#defer-panic-and-recover)
- [Exit and Status Codes](#exit-and-status-codes)
- [Init Functions](#init-functions)

## Overview

- Go is statically typed, high level, memory safe, and garbage collected.
- The language focuses on clean syntax, readability, maintainability, built-in testing, and strong tooling.
- A Go program starts with a package declaration.
- `package main` and `func main()` define the entry point.
- If `main` is missing, the program will not run as an executable.

## Tooling and Execution

- `go build` compiles the program into a binary.
- `go run` is used for rapid development and testing.
- `go run` compiles the source into a temporary binary, runs it, and then removes the temporary artifact.

### Example

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, Go")
}
```

### Output

```text
Hello, Go
```

### Compiler and Runtime

- The compiler translates Go source code into machine code.
- The runtime manages execution after compilation.
- The runtime handles memory allocation, garbage collection, goroutine scheduling, and other execution concerns.
- Go concurrency relies heavily on goroutines, which are lightweight threads managed efficiently by the runtime.
- Even though Go compiles to machine code, the runtime is still needed for memory management and concurrent execution.

## Packages, Imports, and Standard Library

- The standard library is the built-in set of packages that ships with Go.
- These packages provide foundational functionality.
- Example: `import "fmt"` lets you use formatting and printing functions.
- Imports are optimized by the compiler and linker so unused code is excluded from the final binary.
- This is similar to tree shaking in JavaScript.
- Static analysis marks unused code as dead code.
- Named imports allow aliases, for example `import foo "net/http"`.

### Example

```go
package main

import (
  foo "net/http"
  "fmt"
)

func main() {
  fmt.Println(foo.MethodGet)
}
```

### Output

```text
GET
```

## Types, Values, and Naming

- Common Go data types include integers, floats, complex numbers, arrays, strings, constants, structs, pointers, maps, slices, booleans, functions, channels, and JSON-related values.
- If a variable is not initialized, it gets the zero value for its type.
- `:=` is short variable declaration syntax.
- `:=` can be used inside functions, not at package scope.
- At package scope, use `var`.
- Global scope in Go is limited to the package, not the entire program.
- Variables live within their scope, which helps memory efficiency.
- Naming conventions commonly used in Go:
  - PascalCase for exported structs, interfaces, and types.
  - mixedCase for local variables and identifiers.
  - ALL_CAPS is sometimes used for constants.

### Constants

- Use `const` for constants.
- Constants must be initialized so their value is known at compile time.
- Constants cannot use short declaration.

### Overflow

- Overflow happens when a result exceeds the maximum value the type can hold.
- Underflow happens when a result goes below the minimum value the type can hold.

## Arrays, Slices, Maps, and Runes

### Arrays

- Arrays are fixed-size collections of elements of the same type.
- They are stored in contiguous memory.
- Go initializes array elements with zero values.
- Arrays are value types, so assignment copies the array.

### Example

```go
package main

import "fmt"

func main() {
  a := [3]int{1, 2, 3}
  b := a
  b[0] = 99

  fmt.Println(a)
  fmt.Println(b)
}
```

### Output

```text
[1 2 3]
[99 2 3]
```

### Runes

- A rune is a Unicode code point in Go.
- Go uses UTF-8 encoding.

### Maps

- Maps are hash tables in Go.
- They provide average $O(1)$ lookup and average $O(1)$ insertion.
- Create a map with `make(map[string]int)`.
- Reading a missing key returns the zero value.
- Writing to a nil map panics because the internal hash table has not been allocated yet.

### Example

```go
package main

import "fmt"

func main() {
  m := make(map[string]int)
  m["x"] = 10

  v := m["x"]
  fmt.Println(v)
}
```

### Output

```text
10
```

### Map Iteration

- Go intentionally randomizes map iteration order.
- This prevents developers from depending on order.

### Slices

- A slice is a dynamic view over an array.
- Internally, a slice contains a pointer, length, and capacity.
- When capacity is exceeded, Go allocates a new array, copies the old elements, and points the slice to the new storage.
- Length is the number of elements currently in the slice.
- Capacity is how much underlying storage is available before reallocation.

## Functions and Variadic Arguments

- Functions are reusable blocks of code.
- Parameters are copied, so the original passed value is unchanged unless you pass a pointer or reference type.
- Go commonly returns results with an error value, for example `result, err := someFunction()`.
- Errors are values in Go.
- Functions are first-class values.
- You can assign a function to a variable.
- Higher-order functions accept or return functions.
- Closures are anonymous functions that can be assigned or executed immediately.
- Variadic functions accept a variable number of arguments.
- The `...` operator expands a slice when passing it to a variadic parameter.

### Example

```go
package main

import "fmt"

func add(nums ...int) int {
  total := 0
  for _, num := range nums {
    total += num
  }
  return total
}

func main() {
  values := []int{1, 2, 3}
  fmt.Println(add(values...))
}
```

### Output

```text
6
```

## Defer, Panic, and Recover

### Defer

- `defer` means run this later, after the surrounding function finishes.
- Deferred calls are pushed onto a stack and executed in last-in, first-out order.
- Deferred expressions are evaluated when `defer` is encountered, but executed later.
- `defer` is useful for cleanup such as closing database connections or unlocking mutexes.

### Example

```go
package main

import "fmt"

func main() {
  defer fmt.Println("third")
  defer fmt.Println("second")
  fmt.Println("first")
}
```

### Output

```text
first
second
third
```

### Panic

- `panic` stops normal execution immediately.
- When a panic happens, the current function stops, deferred functions execute, and the stack unwinds upward.
- This is called stack unwinding.
- `panic` is for catastrophic or impossible situations, not normal expected errors.

### Example

```go
package main

import "fmt"

func process() {
  fmt.Println("Step 1")
  panic("Something went wrong")
  fmt.Println("Step 2")
}

func main() {
  process()
}
```

### Output

```text
Step 1
panic: Something went wrong
```

### Recover

- `recover` lets you handle a panic inside a deferred function.
- It only works during panic handling inside deferred execution.
- Use it carefully and only where recovery is genuinely appropriate.

### Panic vs Error

- Use `error` for expected issues.
- Use `panic` for unexpected catastrophic issues.
- Most professional Go code uses `if err != nil` instead of `panic`.

## Exit and Status Codes

- `os.Exit()` stops the program immediately.
- Deferred calls do not run when `os.Exit()` is called.
- `os.Exit(0)` means success.
- `os.Exit(1)` means failure or error.
- Status codes matter because operating systems and scripts use them.

### Example

```go
package main

import "os"

func main() {
  os.Exit(1)
}
```

## Init Functions

- `init()` runs before `main()`.
- The Go runtime calls `init()` automatically.
- If package A imports package B, B initializes first.

### Initialization Flow

```text
config init()
service init()
main init()
main()
```

- Dependencies initialize before the code that depends on them.
- Too much logic in `init()` is usually avoided because it hides behavior and makes debugging harder.

## Quick Recall

- Go is compiled, but the runtime still matters.
- Use `make` for maps and slices when you need initialized backing storage.
- Prefer errors for expected failures.
- Use `panic` only for truly exceptional cases.
- Keep `init()` light and predictable.
