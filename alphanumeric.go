//link https://www.codewars.com/kata/526dbd6c8c0eb53254000110/train/go

package main

import "fmt"

func alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}
	// Iterate over each character in the string
	for _, ch := range str {
		if !(ch >= 'a' && ch <= 'z') && !(ch >= 'A' && ch <= 'Z') && !(ch >= '0' && ch <= '9') {
			return false
		}
	}
	// If all characters are letters or digits, the string is valid
	return true
}
func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(alphanumeric(str))
}
