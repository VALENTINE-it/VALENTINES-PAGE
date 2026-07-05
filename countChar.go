/*Count Words in a Sentence
Input: "Go is awesome"
Output: 3
*/

package main

import (
	"fmt"
)
func Countchar (s string) int {
count := 0
inWord := false

	for _, char := range s {
	
		if char != ' ' && char != '\t' && char != '\n' {
			if !inWord {
				count++
				inWord = true
			}
		} else {
			inWord = false
		}
	}
	return count
}


func main() {
	fmt.Println (Countchar("Hello World"))
}