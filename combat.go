
//https://www.codewars.com/kata/586c1cf4b98de0399300001d

package main

import "fmt"

func combat(health, damage float64) float64 {
	if health > damage {
		return health - damage
	}
	return 0
}

func main() {
	fmt.Println(combat(30.43, 50.53))
	fmt.Println(combat(89.43, 50.53))
}
