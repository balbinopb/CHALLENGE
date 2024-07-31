
//https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad

package main

import "fmt"

func Invert(arr []int) []int {
  inverted := make([]int, len(arr))
    for i, num := range arr {
        inverted[i] = -num
    }
    return inverted
}

func main(){

arr:= []int{1, 2, 3, 4, 5}

fmt.Println(Invert(arr))

}