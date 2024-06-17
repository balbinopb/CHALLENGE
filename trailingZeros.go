
//https://www.codewars.com/kata/52f787eb172a8b4ae1000a34/go

package main

import "fmt"

func Zeros(n int) int {
	count := 0
	for i := 5; i <= n; i *= 5 {
		count += n / i
	}
	return count
}

func main() {
	fmt.Println(Zeros(6))
	fmt.Println(Zeros(12))
}
