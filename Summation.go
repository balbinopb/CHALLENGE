package main

import "fmt"

func Summation(n int) int {
	// the sleeper must awaken!
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {

	fmt.Println(Summation(8))

}
