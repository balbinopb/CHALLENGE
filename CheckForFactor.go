
//https://www.codewars.com/kata/55cbc3586671f6aa070000fb

package main

import (
	"fmt"
	)

func CheckForFactor(base int, factor int) bool {
    // your code here
  return base%factor==0
}

func main(){

	fmt.Println(CheckForFactor(2,7))

}