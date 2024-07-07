package main

import"fmt"

func RemoveChar(word string) string {
  if len(word) <= 2 {
    return ""
  }
  return word[1:len(word)-1]
}

func main(){
	fmt.Println(RemoveChar("hello"))
	fmt.Println(RemoveChar("world"))
}