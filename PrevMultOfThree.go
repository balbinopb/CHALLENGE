
//https://www.codewars.com/kata/61123a6f2446320021db987d

package main

import "fmt"

func PrevMultOfThree(n int) interface{} {
  // write your code here
  // your function should return an int or a nil
  for n>1 {
    
    if n%3==0{
      return n
      break
    }
    n =n/10
  }
  return nil
  
}

func main(){

fmt.Println(PrevMultOfThree(36))
fmt.Println(PrevMultOfThree(1244))


}