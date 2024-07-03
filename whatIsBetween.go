package main

import "fmt"


func Between(a, b int) []int {
  arr := []int{}
  
  for i := a; i < b + 1; i++ {
    arr = append(arr, i)
  }
  
  return arr
}

func main(){
	fmt.Println(Between(1,4))
}