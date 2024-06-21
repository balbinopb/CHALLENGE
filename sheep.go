
//https://www.codewars.com/kata/5b077ebdaf15be5c7f000077

package main

import"fmt"

func countSheep(num int) string {
  // Your code here!
  var result string
  for i:=1; i<=num; i++{
    result += fmt.Sprintf("%d sheep...", i)
  }
  return result
}

func main(){
	fmt.Println(countSheep(4))
}