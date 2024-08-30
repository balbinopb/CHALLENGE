
//https://www.codewars.com/kata/5583090cbe83f4fd8c000051

package main

import "fmt"

func Digitize(n int) []int {
  // your code here
  if n>0{
      var arr []int
  for n>0{
    temp:=n%10
    arr=append(arr,temp)
    n=n/10
  }
  return arr
  }
  return []int{0}
}

func main(){

fmt.Println(Digitize(35231))

}