
//https://www.codewars.com/kata/54edbc7200b811e956000556

package main

import "fmt"

func CountSheeps(numbers []bool) int {
  // your code here
  sum:=0
  for i:=0; i<len(numbers); i++{
    if numbers[i]==true{
      sum++
    }
  }
  return sum
}

func main(){

arr:=[]bool{
  true,  true,  true,  false,
  true,  true,  true,  true ,
  true,  false, true,  false,
  true,  false, false, true ,
  true,  true,  true,  true ,
  false, false, true,  true,
	}
fmt.Println(CountSheeps(arr))

}