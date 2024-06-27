
//https://www.codewars.com/kata/5ce9c1000bab0b001134f5af

package main

import "fmt"

func QuarterOf(month int) int {
  // your code here
  var quarter int
  if month>=1 && month<=3{
    quarter=1
  }else if month>=4 && month<=6{
    quarter=2
  }else if month>=7 && month<=9{
    quarter=3
  }else if month>=10 && month<=12{
    quarter=4
  }
  return quarter
}

func main (){
	fmt.Println(QuarterOf(4))
	fmt.Println(QuarterOf(7))
}