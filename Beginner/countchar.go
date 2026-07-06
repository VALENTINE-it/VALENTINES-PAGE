package main 

import "fmt"

func countChar(s string) int {
	count := 0

	for _, i := range s {
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' || i != 'a' && i != 'z' || i != 'A' && i != 'Z'{
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countChar("hello"))
}