
//https://www.codewars.com/kata/580755730b5a77650500010c

package main

import "fmt"

func SortMyString(s string) string {
	var evenChars, oddChars []rune
	for i, char := range s {
		if i%2 == 0 {
			evenChars = append(evenChars, char)
		} else {
			oddChars = append(oddChars, char)
		}
	}
	return string(evenChars) + " " + string(oddChars)
}

func main() {
	str := "CodeWars"
	fmt.Println(SortMyString(str))
}
