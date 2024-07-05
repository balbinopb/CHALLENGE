package main

import "fmt"

func Divisors(n int)int{
  //Put your code here
  sumDivisor:=0
  for i:=1; i<=n; i++{
    if n%i==0{
      sumDivisor++
    }
  }
  return sumDivisor
}

func main(){
	fmt.Println(Divisors(8))
	fmt.Println(Divisors(43))
}