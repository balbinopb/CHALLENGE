
//https://www.codewars.com/kata/5a34b80155519e1a00000009


package main

import "fmt"

func multipleOfIndex (ints []int) []int {
// good luck
  var arr []int
  
  for i:=1; i<len(ints); i++{
    if ints[i]%i==0{
      arr = append(arr,ints[i])
    }
  }
  return arr
}

func main(){

arr:= []int{22, -6, 32, 82, 9, 25}

fmt.Println(multipleOfIndex(arr))

}