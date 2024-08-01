
//https://www.codewars.com/kata/5b180e9fedaa564a7000009a

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func solve(str string) string {
	low := 0
	upp := 0
	for _, char := range str {
		if unicode.IsLower(char) {
			low++
		} else if unicode.IsUpper(char) {
			upp++
		}
	}
	if low >= upp {
		return strings.ToLower(str)
	} else {
		return strings.ToUpper(str)
	}
}

func main(){
	fmt.Println(solve("coDe"))
	fmt.Println(solve("hElLo"))
}