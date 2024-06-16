
//https://www.codewars.com/kata/514b92a657cdc65150000006

package main

import "fmt"

func Multiple3And5(number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		if i%3 == 0 && i%5 == 0 {
			sum += i
		} else if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}
	return sum
}
func main() {
	fmt.Println(Multiple3And5(45))
}
