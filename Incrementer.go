
//https://www.codewars.com/kata/590e03aef55cab099a0002e8

package main

import "fmt"


func Incrementer(n []int) []int {
      // your code here
  for i:=0; i<len(n); i++{
    inc:=n[i]+i+1
    if inc>9 {
      n[i]=inc%10
    }else{
      n[i]=inc
    }
    
  }
  return n
}

func main(){

arr:= []int{1,2,3}

fmt.Println(Incrementer(arr))

}