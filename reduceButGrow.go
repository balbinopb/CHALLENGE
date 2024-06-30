package main

import "fmt"

func Grow(arr []int) int{
  multiply:=1
  for i:=0; i<len(arr); i++{
    multiply *=arr[i]
  }
  return multiply
}

func main(){
	arr:= []int{1, 2, 3, 4}
	fmt.Println(Grow(arr))
}