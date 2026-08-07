package main

import "fmt"

func FirstWord(s string) string {
    for v := 0; v < len(s); v++ {
		if s[v] == ' ' {
			return s[:v] + "\n"
		}
	}
	return s + "\n"
}

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}