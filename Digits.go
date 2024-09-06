
//https://www.codewars.com/kata/58fa273ca6d84c158e000052

package main

import "fmt"

func Digits(n uint64) int {
if n!=0{
    count:=0
for n!=0{
  n /=10
  count++
}
  return count
}
  return 1
}

func main(){

fmt.Println(Digits(12345))


}