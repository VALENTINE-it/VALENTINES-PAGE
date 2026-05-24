package main

import "fmt"

func addition(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func multiplication(a, b int) int {
	return a * b
}

func division(a, b int) int {
	return a / b
}

func square(a int) int {
	return a * a
}

func main() {
	fmt.Println(addition(5, 3))
	fmt.Println(subtraction(5, 3))
	fmt.Println(multiplication(5, 3))
	fmt.Println(division(5, 3))
	fmt.Println(square(5))
}
