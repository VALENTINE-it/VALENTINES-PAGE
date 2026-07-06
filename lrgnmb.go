package main

import "fmt"

func lrgnmb(s []int) int {
	largest := 0

	for _, num := range s {
		if num > largest {
			largest = num
		}
	}
	return largest
}

func main() {
	arr := []int{3,6,1,8,4,0,2,5}
	fmt.Println(lrgnmb(arr))
}