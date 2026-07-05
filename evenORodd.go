/*Even or Odd Checker
Input a number
Print whether it’s even or odd
*/	

package main

import "fmt"

func main() {
	var number int

	fmt.Println("Please enter number:")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println ("The number is Even")
	}else {
		fmt.Println("The number is odd")
	}
}