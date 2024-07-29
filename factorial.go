
//https://www.codewars.com/kata/57a049e253ba33ac5e000212

package main

import "fmt"

func Factorial(n int) int {
	// put your code here
  if n==0{
    return 1
  }else {
    return n*Factorial(n-1)
  }
}

func main(){

fmt.Println(Factorial(7))
fmt.Println(Factorial(4))

}
