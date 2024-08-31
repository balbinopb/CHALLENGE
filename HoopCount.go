
//https://www.codewars.com/kata/55cb632c1a5d7b3ad0000145

package main

import "fmt"

func HoopCount(n int) string {
  // your code goes here
  if n>9{
    return "Great, now move on to tricks"
  }
	return "Keep at it until you get it"
}

func main(){

fmt.Println(HoopCount(10))

}