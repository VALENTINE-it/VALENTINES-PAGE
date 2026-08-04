/* 
Given an array:

Find the largest element
Find the smallest element
Calculate the average
*/

package main

import "fmt"

func main() {

	numbers := [4]int{2, 4, 8, 9}

	largest := numbers[0]
	smallest := numbers [0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > largest {
			largest = numbers[i]
		}
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}

	sum := numbers[0] + numbers[1] + numbers[2] + numbers[3]

	average := float64(sum) / 4

	fmt.Println("Largest", largest)
	fmt.Println("Smallest", smallest)
	fmt.Println("Average", average)
}