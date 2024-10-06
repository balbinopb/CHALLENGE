
//https://www.codewars.com/kata/55f2b110f61eb01779000053

package main

import "fmt"


func GetSum(a, b int) int {
  if a < b {
    a, b = b, a
  }
  
  var sum int
  
  for i := a; i >= b; i-- {
    sum += i
  }
  
  return sum
}

func main(){

fmt.Println(GetSum(1, 2))

}