
//https://www.codewars.com/kata/51f2d1cafc9c0f745c00037d

package main

import "fmt"

func solution(str, ending string) bool {
	// Your code here!
	if len(ending) > len(str) {
		return false
	}
	return str[len(str)-len(ending):] == ending

}
func main() {

	var start, end string
	start = "abcd"
	end = "d"
	fmt.Println(solution(start, end))
}
