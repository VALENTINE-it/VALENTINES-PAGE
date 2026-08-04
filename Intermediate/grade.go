package main

import "fmt"

func main() {

	var grade string

	fmt.Println("Please Enter Your Grade:")
	fmt.Scanln(&grade)

	if grade == "A" {
		fmt.Println("Passed")
	} else if grade == "B" {
		fmt.Println("Passed")
	}else if grade == "C" {
		fmt.Println("Above Average")
	}else if grade == "D" {
		fmt.Println("Failed")
	}else {
		fmt.Println("Invalid Grade")
	}

} 
