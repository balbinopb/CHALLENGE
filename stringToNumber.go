
//https://www.codewars.com/kata/544675c6f971f7399a000e79

package main

import (
	"fmt"
	"strconv"
)

func StringToNumber(str string) int {
	// Convert the string to an integer
	number, eror := strconv.Atoi(str)
	if eror != nil {
		return 0
	}
	return number
}

func main(){
	fmt.Println(StringToNumber("1234"))
	fmt.Println(StringToNumber("605"))
}