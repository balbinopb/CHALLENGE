package main

import"fmt"

func Derive(coefficient, exponent int) string {
  coefficient=coefficient*exponent
  exponent=exponent-1
  return fmt.Sprintf("%dx^%d", coefficient, exponent)
  //return "{coefficient}x^{exponent}"
}

func main() {
  fmt.Println(Derive(7,8))
  fmt.Println(Derive(5,9))
}