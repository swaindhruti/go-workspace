# Go Intermediate Notes

## Index
- [Closures](#closures)
- [Recursion](#recursion)
- [Pointers](#pointers)
- [Strings and Runes](#strings-and-runes)
- [fmt Package](#fmt-package)
- [Structs](#structs)
- [Methods](#methods)
- [Interfaces](#interfaces)
- [Quick Recall](#quick-recall)

## Closures
- A closure is a function that remembers variables from outside itself.
- The remembered data is closure state.
- In Go, a closure means function plus remembered variables.
- Closures capture variables, not values.
- That is the core behavior behind persistent state between calls.
- `adder()` is only executed once when the closure is created.
- The variable such as `i` is initialized once and then kept alive by the closure.
- This is why `i` does not reset to zero on later calls.
- Closures allow private persistent state without global variables.
- Each new closure gets its own state.
- The garbage collector cannot remove captured variables while the closure still references them.

### Simple Memory Model
```text
Normal function
function runs
-> variables created
-> function ends
-> variables destroyed

Closure
function runs
-> variables created
-> closure captures variables
-> function ends
-> variables kept alive
```

### Example
```go
package main

import "fmt"

func adder() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    a := adder()
    b := adder()

    fmt.Println(a())
    fmt.Println(a())
    fmt.Println(b())
}
```

### Output
```text
1
2
1
```

## Recursion
- Recursion is a function solving a problem by asking a smaller version of itself to solve the problem.
- It is used in backend engineering for file traversal, JSON processing, and tree structures.
- Memoization means storing already calculated answers.

### Recursion Template
```go
func solve(input Type) ReturnType {
    // Base Case
    if condition {
        return answer
    }

    // Recursive Case
    return solve(smallerProblem)
}
```

## Pointers
- A variable stores a value in memory.
- A pointer stores the address of that value.
- `&` means “give me the address.”
- `*` means “give me the value at that address.”
- This is the first thing to understand about pointers.
- Nil pointers have no assigned target.
- A pointer can exist while still pointing nowhere.

### Example
```go
package main

import "fmt"

func main() {
    a := 10
    ptr := &a

    fmt.Println(ptr)
    fmt.Println(*ptr)
}
```

### Output
```text
0xc0000120b0
10
```

### Nil Pointer Example
```go
package main

import "fmt"

func main() {
    var ptr *int
    fmt.Println(ptr)
}
```

### Output
```text
<nil>
```

## Strings and Runes
- A string in Go is a read-only sequence of bytes.
- Strings are immutable after creation.
- You cannot modify a string by indexing into it.
- This makes strings safe for concurrent systems.
- Backticks create raw strings.
- Go does not interpret escape sequences inside raw strings.
- `len()` returns the number of bytes, not the number of characters.
- Strings behave like arrays of bytes when indexed.
- `msg[0]` returns a byte value, not a character.
- For example, `H` is represented by byte value `72`.
- A rune is a Unicode code point.
- In Go, `rune` is an alias for `int32`.
- `utf8.RuneCountInString` counts Unicode characters rather than bytes.
- This matters for multibyte characters such as emoji.

### String Example
```go
package main

import "fmt"

func main() {
    msg := "Hello"
    fmt.Println(len(msg))
    fmt.Printf("%c\n", msg[0])
}
```

### Output
```text
5
H
```

### Rune Count Example
```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    s := "😊"
    fmt.Println(len(s))
    fmt.Println(utf8.RuneCountInString(s))
}
```

### Output
```text
4
1
```

## fmt Package
- `fmt` is a communication package for formatting and input/output.
- `Print` prints values.
- `Sprint` returns a string.
- `Scan()` reads input.
- `Scanln()` reads input and stops at a newline.
- `Scan()` needs addresses because it modifies the original variable.
- `fmt.Errorf()` creates formatted errors with dynamic values.
- Real backend code uses `fmt.Errorf()` constantly.
- `%w` is used for wrapping errors.

### Scan Example
```go
package main

import "fmt"

func main() {
    var name string
    fmt.Scan(&name)
    fmt.Println(name)
}
```

### Output
```text
John
```

### Errorf Example
```go
package main

import "fmt"

func main() {
    id := 123
    err := fmt.Errorf("user %d not found", id)
    fmt.Println(err)
}
```

### Output
```text
user 123 not found
```

### Wrapped Error Example
```go
package main

import "fmt"

func main() {
    err := fmt.Errorf("failed to create user: %w", fmt.Errorf("database unavailable"))
    fmt.Println(err)
}
```

### Output
```text
failed to create user: database unavailable
```

## Structs
- Structs are user-defined composite types.

### Mental Model
- Data is stored in structs.
- Behavior is attached through methods.

### Example
```go
type Person struct {
    Name string
    Age  int
}
```

### Typical Behavior for Person
- `FullName()`
- `GrowOlder()`
- `CanVote()`

### Supporting Notes
- Structs group related fields into one type.
- Struct values can be created using literals.
- Fields with uppercase names are exported outside the package.
- Fields with lowercase names are package-private.

## Methods
- Methods attach behavior to a type.
- Methods belong to types, not only to structs.
- Methods make code more organized than standalone helper functions.

### Why Methods Over Normal Functions
```go
func Area(r Rectangle) float64 {
    return r.Length * r.Width
}
```

```go
func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}
```

- Both work, but method syntax is cleaner in usage: `rect.Area()`.
- Mental model: Rectangle knows how to calculate area.

### Receiver Syntax
```go
func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}
```

- `(r Rectangle)` is called the receiver.
- Value receiver is usually used for read-only behavior.
- Pointer receiver is usually used when data must be modified or when copying should be avoided.

### Value vs Pointer Receiver
```go
func (p Person) CanVote() bool {
    return p.Age >= 18
}

func (p *Person) GrowOlder() {
    p.Age++
}
```

### Method Promotion via Embedding
- Method promotion is a key Go concept.
- It happens when one type embeds another type anonymously.

```go
type Rectangle struct {
    Length float64
    Width  float64
}

func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}

type Shape struct {
    Rectangle
}
```

- Without promotion: `shape.Rectangle.Area()`.
- With promotion: `shape.Area()`.

### Why This Matters
- Go favors composition over inheritance.
- Java style is inheritance-heavy.
- Go style is embedding plus promotion.

## Interfaces
- An interface defines what a type can do, not what a type is.
- Interfaces declare behavior using method signatures.
- Go uses implicit interface satisfaction.
- No `implements` keyword is required.

### Example: Geometry Interface
```go
type Geometry interface {
    Area() float64
    Perim() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64  { return 3.14 * c.Radius * c.Radius }
func (c Circle) Perim() float64 { return 2 * 3.14 * c.Radius }
```

- `Circle` satisfies `Geometry` automatically because required methods exist.

### Why Measure Works
```go
func Measure(g Geometry) {
    fmt.Println(g.Area())
    fmt.Println(g.Perim())
}
```

- The function accepts `Geometry`, not a specific concrete type.
- `Measure(rect)` works if `rect` implements the methods.
- `Measure(circle)` works if `circle` implements the methods.
- This is polymorphism: one interface, many implementations.

### Method Sets and Interface Satisfaction
- If interface methods are implemented with value receivers, both value and pointer can satisfy the interface.
- If implemented only with pointer receivers, only the pointer type satisfies the interface.

```go
type Speaker interface {
    Speak()
}

type User struct{}

func (u *User) Speak() {}

// var s Speaker = User{}   // compile error
// var s Speaker = &User{}  // works
```

### Type Switch
- `switch v := i.(type)` is a type switch.
- It is useful when a function accepts `any` but behavior depends on runtime type.

```go
func PrintType(v any) {
    switch v := v.(type) {
    case int:
        fmt.Println("int", v)
    case string:
        fmt.Println("string", v)
    default:
        fmt.Println("unknown")
    }
}
```

### Output
```text
PrintType(10)     -> int 10
PrintType("John") -> string John
```

### FAQ Notes
- Does Go have inheritance? No. It uses interfaces, embedding, and composition.
- Explicit implementation? No. Satisfaction is implicit.
- Can an interface have data fields? No. Methods only.
- Can a type have extra methods? Yes, as long as required methods exist.

### Empty Interface
- `any` is an alias for `interface{}`.
- It can hold values of any type.

## Quick Recall
- Struct stores data.
- Method defines behavior on a type.
- Interface defines a behavioral contract.
- A type must implement all interface methods to satisfy the interface.
- Go uses implicit satisfaction with no `implements` keyword.

## Generics

### Why Generics Exist

Without generics, you write one function per type:

```go
func PrintInts(nums []int)
func PrintStrings(nums []string)
func PrintBools(nums []bool)
```

Same logic repeated. With generics, one function works for many types:

```go
func Print[T any](items []T) {
    for _, v := range items {
        fmt.Println(v)
    }
}
```

Usage:

```go
Print([]int{1, 2, 3})
Print([]string{"A", "B"})
Print([]bool{true, false})
```

**Generics allow writing code that works with multiple types safely.**

| Without Generics | With Generics |
|---|---|
| One Function, One Type | One Function, Many Types |

### Type Parameters [T]

`[T]` declares a placeholder type. It works like parameters in a function:

```go
func Add(a int, b int)  // a, b are placeholder values
func Print[T any](items []T) // T is a placeholder type
```

Go fills in `T` at the call site:

```
T → int     when calling Print([]int{1, 2})
T → string  when calling Print([]string{"A"})
T → User    when calling Print([]User{})
```

Inference is automatic -- you rarely need to specify `T` explicitly.

### The `any` Constraint

`any` is an alias for `interface{}`. `[T any]` means T can be any type, with no restrictions.

```go
[T any]  // T can be anything -- int, string, struct, pointer, etc.
```

### Generic Types and Methods

Types can also be generic:

```go
type Box[T any] struct {
    Value T
}

func (b Box[T]) Get() T {
    return b.Value
}
```

Usage:

```go
box := Box[int]{10}
fmt.Println(box.Get()) // 10
```

The type parameter `[T]` must be repeated on the type definition and any methods attached to it.

### Constraints

`any` is too permissive for operations like `+`:

```go
func Add[T any](a, b T) T {
    return a + b // compile error
}
```

The compiler rejects this because `any` allows types that do not support `+`.

Use a **constraint** to restrict allowed types:

```go
type Number interface {
    ~int | ~float64
}

func Add[T Number](a, b T) T {
    return a + b
}
```

Now `Add(10, 20)` and `Add(1.5, 2.5)` work, but `Add(User{}, User{})` is a compile error.

Constraints use the `interface` syntax with type unions (`|`).

### The `~` Symbol (Underlying Type)

`~int` matches `int` and any type whose underlying type is `int`:

```go
type Age int  // underlying type is int

type Number interface {
    ~int
}

func Accept[T Number](v T) {}

Accept(10)    // int
Accept(Age(5)) // Age -- allowed because ~int matches the underlying type
```

Without `~`, only `int` would match, not `Age`.

### Interfaces vs Generics

| Interfaces | Generics |
|---|---|
| Behavior abstraction | Type abstraction |
| "What a type can do" | "What types are allowed" |
| Method signatures | Type constraints |