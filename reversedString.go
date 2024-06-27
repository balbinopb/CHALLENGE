
//https://www.codewars.com/kata/5168bb5dfe9a00b126000018

package main

import"fmt"

func Solution(word string) string {
  var result string
  lenth:=len(word)-1
  for i:=lenth;i>=0;i--{
    result +=string(word[i])
  }
  return result
}

func main(){
	fmt.Println(Solution("test"))
	fmt.Println(Solution("hello"))
}