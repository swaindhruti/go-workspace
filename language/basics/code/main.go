package main // entry point of the program

import (
	"fmt"
	"os"
) // importing package

func hello_world() {
	fmt.Println("Hello, World!")
}

// variable declaration and initialization
func variable_declaration() {
	var age int
	age = 25
	var surname = "Swain" // we can ignore the type as it can be inferred from the value (Type inference)
	name := "Dhrutinandan"
	fmt.Printf("My name is %s %s and I am %d years old.\n", name, surname, age)
}

// performing basic arithmetic operations
func arithmetic_operations() {
	var a, b int = 10, 5
	fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Subtraction: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, a*b)
	fmt.Printf("Division: %d / %d = %d\n", a, b, a/b)
	fmt.Printf("Remainder: %d %% %d = %d\n", a, b, a%b)

	var maxInt int64 = 9223372036854775807
	fmt.Printf("Max int64 value: %d\n", maxInt)
	maxInt++ // this will cause overflow
	fmt.Printf("After overflow: %d\n", maxInt)

	var minInt int64 = -9223372036854775808
	fmt.Printf("Min int64 value: %d\n", minInt)
	minInt-- // this will cause underflow
	fmt.Printf("After underflow: %d\n", minInt)

}

// loop demonstration
func loop_demo() {
	fmt.Println("For loop:")
	for i := range 5 { // simple iteration over range
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Println("Break and Continue demonstration:")
	for i := range 10 { // demonstration of continue and break
		if i%2 == 0 {
			continue // will skip and will go with next iteration
		}
		if i == 5 {
			break // will exit the loop
		}
		fmt.Printf("Odd no. %d\n", i)
	}

	fmt.Println("While loop using for loop:")
	// while loop using for and break
	count := 0
	for {
		if count >= 5 {
			break
		}
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Println()

	// while loop using condition in for loop
	fmt.Println("While loop using condition in for loop:")
	count = 0
	for count < 5 {
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Println()

}

// Operators in Go
func operators_demo() {
	// Logical operators
	a, b := true, false
	fmt.Printf("a AND b: %t\n", a && b)
	fmt.Printf("a OR b: %t\n", a || b)
	fmt.Printf("NOT a: %t\n", !a)

	// Comparison operators
	x, y := 10, 20
	fmt.Printf("x == y: %t\n", x == y)
	fmt.Printf("x != y: %t\n", x != y)
	fmt.Printf("x > y: %t\n", x > y)
	fmt.Printf("x < y: %t\n", x < y)
	fmt.Printf("x >= y: %t\n", x >= y)
	fmt.Printf("x <= y: %t\n", x <= y)

	// Bitwise operators
	c, d := 5, 3                               // 5 = 0101, 3 = 0011 in binary
	fmt.Printf("c & d (AND): %d\n", c&d)       // 1 (0001)
	fmt.Printf("c | d (OR): %d\n", c|d)        // 7 (0111)
	fmt.Printf("c ^ d (XOR): %d\n", c^d)       // 6 (0110)
	fmt.Printf("c &^ d (AND NOT): %d\n", c&^d) // 4 (0100)
}

// Control flow demonstration
func control_flow_demo() {
	// if-else demonstration
	num := 10
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	// switch case demonstration
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the week")
	default:
		fmt.Println("Midweek")
	}

	// else if demonstration
	score := 85
	if score >= 90 { // also known as short circuit branching as it will not check the next condition if the first condition is true
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
}

// Array, Slice, Map, Range demonstration
func array_slice_map_demo() {
	// Array demonstration
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v\n", arr)

	// Pointer to array
	ptr_array := &arr
	ptr_array[0] = 10 // modifying the array through pointer
	fmt.Printf("Pointer to array: %v\n", ptr_array)

	// Slice demonstration
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v\n", slice)

	// Map demonstration
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Printf("Map: %v\n", m)

	// Range demonstration
	fmt.Println("Range over array:")
	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// Range over map
	fmt.Println("Range over map:")
	for k, v := range m {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	// Range over string
	str := "Hello"
	fmt.Println("Range over string:")
	for i, r := range str { // r is of type rune which is an alias for int32 and represents a Unicode code point
		fmt.Printf("Index: %d, Value: %c\n", i, r)
	}

	// Slice is a reference type and it points to an underlying array. When we create a slice, it creates a reference to the array and any changes made to the slice will affect the underlying array. This is why when we modify the slice, the original array also gets modified.
	arr2 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr2[1:4]
	slice2[3] = 10 // modifying the slice will modify the underlying array
	fmt.Printf("Original array after modifying slice: %v\n", arr2)
	slice2 = append(slice2, 20) // appending to slice will create a new underlying array if the capacity is exceeded
	fmt.Printf("Slice after appending: %v\n", slice2)
	fmt.Printf("Original array after appending to slice: %v\n", arr2) // original array remains unchanged as a new underlying array is created for the slice
}

// Defer demonstration
func defer_demo() {
	defer fmt.Println("This will be printed at the end of the function execution") // defer statements are executed in LIFO order
	defer fmt.Println("This will be printed fourth")
	defer fmt.Println("This will be printed third")
	defer fmt.Println("This will be printed second")
	fmt.Println("This will be printed first and defer follows LIFO order")

	x := 10
	defer fmt.Println("Value of x at the time of defer: ", x) // defer captures the value of x at the time of defer statement
	x = 20
	fmt.Println("Value of x after modification: ", x) // this will not affect the deferred statement as it has already captured the value of x
}

// Recover demonstration
func panic_recover_demo() {

	defer func() { // recover must be called inside a deferred function to catch the panic and allow the program to continue execution
		r := recover()

		if r != nil {
			fmt.Println("Recovered from panic: ", r) // recover will catch the panic and allow the program to continue execution
		}
	}()

	fmt.Println("This will be printed first")
	panic("This is a panic") // this panic will be caught by the recover function in the deferred function
	// fmt.Println("This will not be printed due to panic") -> this will not be executed as the panic will stop the normal execution of the program and will start unwinding the stack
}

// os.Exit demonstration
func os_exit_demo() {
	fmt.Println("This will be printed first")
	os.Exit(1) // os.Exit will immediately terminate the program without executing deferred statements and without handling any panics
	// fmt.Println("This will not be printed due to os.Exit") -> this will not be executed as os.Exit will immediately terminate the program
}

// main function is the entry point of the program
func main() {
	var choice int
	fmt.Println("Enter\n1 for Hello World\n2 for Variable Declaration\n3 for arithmetic operations\n4 for loop demonstration\n5 for operators demonstration\n6 for control flow demonstration\n7 for array, slice, map and range demonstration\n8 for defer, panic and recover demonstration")
	fmt.Scan(&choice)
	fmt.Println("****************************************************************")

	switch choice { // switch case for user input
	case 1:
		hello_world()
	case 2:
		variable_declaration()
	case 3:
		arithmetic_operations()
	case 4:
		loop_demo()
	case 5:
		operators_demo()
	case 6:
		control_flow_demo()
	case 7:
		array_slice_map_demo()
	case 8:
		defer_demo()
		panic_recover_demo()
	default:
		fmt.Println("Invalid choice")
	}
}

// init function is a special function in Go that is executed before the main function and it is used for initialization purposes. We can have multiple init functions in a package and they will be executed in the order they are defined.
func init() {
	fmt.Println("This will be printed before the main function execution") // we can have multiple init functions and they will be executed in the order they are defined
}
