package main

import "fmt"

func FindMultiples(integer, limit int) []int {
  // Your code here!
  var arr []int
  for i:=integer; i<=limit; i +=integer{
    arr=append(arr,i)
  }
  return arr
}

func main(){


fmt.Println(FindMultiples(2,6))

}