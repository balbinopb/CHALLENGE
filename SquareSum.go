//https://www.codewars.com/kata/515e271a311df0350d00000f

package main

import "fmt"

func SquareSum(numbers []int) int {
    // your code here
  sum:=0
  for i:=0; i<len(numbers); i++{
    sum += numbers[i]*numbers[i]
  }
  return sum
}
func main(){

arr:= []int{1,2,2}

fmt.Println(SquareSum(arr))

}