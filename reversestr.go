package main

import "fmt"

func reverseStr(s string) string {
	result := ""

	for i := len(s)-1; i >=0; i-- {
		result += string(s[i])
	}  
	return result
}

func main() {
	fmt.Println(reverseStr("hellowearehuman"))
}