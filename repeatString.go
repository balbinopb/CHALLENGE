
//https://www.codewars.com/kata/57a0e5c372292dd76d000d7e

package main

import "fmt"

func RepeatStr(repetitions int, value string) string {
  sumStr:=value
  for i:=1; i<repetitions;i++{
    sumStr +=value
  }
  return sumStr
}

func main(){

	fmt.Println(RepeatStr(4,"hello"))
}