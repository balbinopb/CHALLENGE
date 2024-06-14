
//https://www.codewars.com/kata/5839edaa6754d6fec10000a2

package main

import "fmt"

func FindMissingLetter(chars []rune) rune {
	for i := 0; i < len(chars)-1; i++ {
		if chars[i+1] != chars[i]+1 {
			return chars[i] + 1
		}
	}
	return 0
}
func main() {
	chars := []rune{'a', 'b', 'c', 'd', 'f'}
	fmt.Println(FindMissingLetter(chars))
}
