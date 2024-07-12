//https://www.codewars.com/kata/5a03b3f6a1c9040084001765

package main

import "fmt"

func Angle(n int) int {
	// your code here
	//Sum of internal angles=(n−2)×180
	return (n - 2) * 180
}

func main() {
	fmt.Println(Angle(4))
	fmt.Println(Angle(7))
}
