package main

//https://www.codewars.com/kata/583f158ea20cfcbeb400000a


import "fmt"

func Arithmetic(a int, b int, operator string) int{
  //your code here
  var result int
  switch operator  {
    case "add":
      result=a+b
    case "subtract":
      result=a-b
    case "multiply":
      result=a*b
    case "divide":
      result=a/b
}
  return result
}

func main(){

fmt.Println(Arithmetic(4,6,"add"))
fmt.Println(Arithmetic(4,6,"multiply"))


}