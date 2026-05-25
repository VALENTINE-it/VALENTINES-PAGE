package main

import "fmt"

func main() {
	var name string
	var number int
	var location string
	var email string

	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)

	fmt.Println("Please enter your Email Address:")
	fmt.Scanln(&email)

	fmt.Println("Please enter your phone number:")
	fmt.Scanln(&number)

	fmt.Println("Please enter your location:")
	fmt.Scanln(&location)

	fmt.Println("Please confirm your details:")
	fmt.Println("Your Name is:", name)
	fmt.Println("Yor Email Address is:", email)
	fmt.Println("Your Phone Number is:", number)
	fmt.Println("Your Location is:", location)
}
