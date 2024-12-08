// Package main is the starting point of the Go program.
package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Constants in Go
const Pi = 3.14

// Variables in Go
var (
	name   string = "Saurabh"
	age    int    = 29
	active bool   = true
)

// Structs in Go: Define a custom type `Person`.
type Person struct {
	firstName string
	lastName  string
	age       int
}

// Method on the Person struct
// Methods are functions with a receiver, in this case, a pointer to Person.
func (p *Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// Interfaces in Go
// An interface defines a set of method signatures.
type Greetable interface {
	greet() string
}

// Implement the Greetable interface for the Person struct
func (p *Person) greet() string {
	return "Hello, my name is " + p.fullName()
}

// Functions in Go
// A simple function that adds two numbers and returns the result.
func add(a int, b int) int {
	return a + b
}

// A function that demonstrates error handling
func divide(a, b int) (int, error) {
	if b == 0 {
		// Return an error using the errors package
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Function demonstrating flow control (for loop)
func printNumbers(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}

// Go routines and concurrency
// A simple function that prints a message after a delay
func delayedPrint(message string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done
	time.Sleep(delay)
	fmt.Println(message)
}

func main() {
	// Basic types in Go
	var a int = 10
	var b float64 = 20.5
	var c string = "Hello, Go!"
	var d bool = true

	// Print the basic types
	fmt.Println("Integer:", a)
	fmt.Println("Float:", b)
	fmt.Println("String:", c)
	fmt.Println("Boolean:", d)

	// Working with structs
	p := Person{firstName: "John", lastName: "Doe", age: 30}
	fmt.Println("Person's full name:", p.fullName())

	// Working with interfaces
	var greeter Greetable = &p
	fmt.Println(greeter.greet())

	// Working with functions
	sum := add(10, 20)
	fmt.Println("Sum:", sum)

	// Error handling in functions
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", result)
	}

	// Flow control with for loop
	fmt.Println("Printing numbers from 1 to 5:")
	printNumbers(5)

	// Demonstrate concurrency with goroutines
	var wg sync.WaitGroup
	wg.Add(2) // Add two goroutines to the WaitGroup

	go delayedPrint("First message", 2*time.Second, &wg)
	go delayedPrint("Second message", 1*time.Second, &wg)

	wg.Wait() // Wait for all goroutines to finish

	// Example of using a constant
	fmt.Println("Value of Pi:", Pi)

	// Demonstrating the use of a package-level variable
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Active?", active)
}
