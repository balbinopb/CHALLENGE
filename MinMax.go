
//https://www.codewars.com/kata/559590633066759614000063

package main

import "fmt"


func MinMax(arr []int) [2]int {
  max:=arr[0]
  min:=arr[0]
  for i:=1; i<len(arr); i++{
    if arr[i]>max{
      max=arr[i]
    }
    if arr[i]<min{
      min=arr[i]
    }
  }
  return [2]int{min,max}
}
func main(){

arr:= []int{1,2,3,4,5}
fmt.Println(MinMax(arr))

}