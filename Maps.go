
//https://www.codewars.com/kata/57f781872e3d8ca2a000007e

package main

import "fmt"

func Maps(x []int) []int {
  result := make([]int, len(x))

  for i := 0; i < len(x); i++ {
    result[i] = x[i] * 2
  }

  return result
}

func main(){

arr := []int{1,2,3}

fmt.Println(Maps(arr))

}