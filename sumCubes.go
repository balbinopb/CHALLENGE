
//https://www.codewars.com/kata/59a8570b570190d313000037/go

package main

import "fmt"

func SumCubes(n int) int {
	// your code here
  sum:=0
  for i:=1; i<=n; i++{
    sum +=i*i*i
  }
  return sum
}

func main(){

fmt.Println(SumCubes(4))
fmt.Println(SumCubes(2))

}