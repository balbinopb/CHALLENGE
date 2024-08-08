
//https://www.codewars.com/kata/576bb71bbbcf0951d5000044

package main

import "fmt"

func CountPositivesSumNegatives(numbers []int) []int {
  var res []int
  sumNegative:=0
  countPositive:=0
  for i:=0; i<len(numbers); i++{
    if numbers[i]>0{
      countPositive++
    }else{
      sumNegative+=numbers[i]
    }
  }
  res=append(res,countPositive)
  res=append(res,sumNegative)
  
  return res
 // your code here
}

func main(){

arr:= []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}

fmt.Println(CountPositivesSumNegatives(arr))

}