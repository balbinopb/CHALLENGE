
//https://www.codewars.com/kata/5715eaedb436cf5606000381

package main

import "fmt"

func PositiveSum(numbers []int) int {
  sum:=0
  for i:=0; i<len(numbers); i++{
     if numbers[i]>0{
       sum+=numbers[i]
     }    
  }
  return sum
}

func main(){

arr:= []int{1,-4,7,12}

fmt.Println(PositiveSum(arr))

}