
//https://www.codewars.com/kata/565f5825379664a26b00007c/train/go

package main

import "fmt"

func GetSize(w,h,d int) [2]int {
    // your code here
  return [2]int{(2*w*h+2*w*d+2*h*d),(w*h*d)}
} 

func main(){

fmt.Println(GetSize(4, 2, 6))

}