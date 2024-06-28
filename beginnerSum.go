
//https://www.codewars.com/kata/55f2b110f61eb01779000053

package main

import "fmt"

func GetSum(a, b int) int {
  sum:=0
  if a > b{
    for i:=b; i<=a; i++{
      sum +=i
    }
  }else if b>a {
    for i:=a; i<=b; i++{
      sum +=i
    }
  }else {
    sum=a
  }
  return sum
}

func main(){
	fmt.Println(GetSum(1,2))
}