
//https://www.codewars.com/kata/59ca8246d751df55cc00014c

package main

import "fmt"

func Hero(bullets, dragons int) bool {
  // your code
  return bullets>= 2*dragons
}

func main(){

fmt.Println(Hero(10, 5))

}