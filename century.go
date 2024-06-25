package main

import"fmt"

func century(year int) int {
 // Finish this :)
  var century int
  if year % 100 !=0{
    century = year/100 + 1
  }else {
    century =year/100
  }
 return century
}

func main(){
	fmt.Println(century(89))
}