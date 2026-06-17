package main

import (
	"fmt"
	"unicode/utf8"
)

// Closure is a function that captures the variables from its surrounding scope. It allows the function to access and manipulate those variables even after the outer function has finished executing.
func closure_demo() {

	x := 10

	closure := func() {
		println(x)
	}

	closure()
}

// Recursion is a programming technique where a function calls itself in order to solve a problem. It typically involves a base case that stops the recursion and a recursive case that breaks the problem into smaller subproblems.
func recursion_demo(n int) int {
	if n == 0 {
		return 1
	}

	result := n * recursion_demo(n-1)
	fmt.Printf("Factorial of %d is %d\n", n, result)

	return result
}

// Pointers are variables that store the memory address of another variable. They allow you to indirectly access and manipulate the value stored at that memory address. Pointers are commonly used in programming languages like C and C++ to enable dynamic memory allocation and efficient data manipulation.
func pointer_demo() {
	var x int = 10
	var p *int = &x // p is a pointer to x

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Value of p (address of x): %p\n", p)
	fmt.Printf("Value pointed to by p (value of x): %d\n", *p)

	*p = 20 // Modifying the value at the address pointed to by p
	fmt.Printf("New value of x after modification through pointer: %d\n", x)
}

// Strings and runes in Go are used to represent text. A string is a sequence of bytes, while a rune is a Unicode code point that represents a single character. Runes are used to handle characters that may require more than one byte, such as those in non-ASCII languages. In Go, you can convert a string to a slice of runes to work with individual characters, and you can also use the built-in functions to manipulate strings and runes effectively.
func string_and_rune_demo() {
	str := "Hello, 世界"
	fmt.Printf("String: %s\n", str)

	// Convert string to slice of runes
	runes := []rune(str)
	fmt.Printf("Runes: %v\n", runes)

	// Accessing individual characters
	for i, r := range runes {
		fmt.Printf("Character at index %d: %c\n", i, r)
	}

	// Length of string in bytes and runes
	fmt.Printf("Length of string in bytes: %d\n", len(str))
	fmt.Printf("Length of string in runes: %d\n", len(runes))

	// Rune Count
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(str))
}

// Structs are composite data types that group together variables (fields) under a single name. They allow you to create custom data types that can represent real-world entities or concepts. Structs are commonly used in Go to define complex data structures and organize related data. You can define methods on structs to provide behavior and functionality specific to that struct type.
func struct_demo() {
	type Address struct {
		Street string
		City   string
		State  string
	}

	type Person struct {
		Name    string
		Age     int
		Address Address // Struct Embedding: Address is embedded within Person, allowing access to its fields directly through the Person struct. This is composition, as Person is composed of Address. Go prefers composition over inheritance, allowing for more flexible and modular code design.
	}

	address := Address{
		Street: "123 Main St",
		City:   "Anytown",
		State:  "CA",
	}

	person := Person{
		Name:    "John Doe",
		Age:     30,
		Address: address,
	}

	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Address: %s, %s, %s\n", person.Address.Street, person.Address.City, person.Address.State)
}

func main() {
	var choice int
	fmt.Println("Enter\n1 for closure demo\n2 for recursion demo\n3 for pointer demo\n4 for string and rune demo\n5 for struct demo")
	fmt.Scan(&choice)
	fmt.Println("****************************************************************")

	switch choice { // switch case for user input
	case 1:
		closure_demo()
	case 2:
		recursion_demo(5)
	case 3:
		pointer_demo()
	case 4:
		string_and_rune_demo()
	case 5:
		struct_demo()
	default:
		fmt.Println("Invalid choice")
	}
}
