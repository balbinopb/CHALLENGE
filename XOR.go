package main

import "fmt"

func Xor(a bool, b bool) bool {
  // your code here:
  return a != b
}

func main(){
	fmt.Println(Xor(true,true))
	fmt.Println(Xor(true,false))
}