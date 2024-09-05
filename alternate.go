
//https://www.codewars.com/kata/62a611067274990047f431a8

package main

import "fmt"

func Alternate(n int, firstValue string, secondValue string) []string {
  var arr []string
  if n>0{
    for i:=0;i<n; i++{
      if (i+1)%2 !=0{
        arr=append(arr,firstValue)
      }else{
        arr=append(arr,secondValue)
      }
      
    }
    return arr
  } 
  return []string{}
}


func main(){

fmt.Println(Alternate(10, "blue", "red"))


}