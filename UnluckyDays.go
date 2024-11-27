
//https://www.codewars.com/kata/56eb0be52caf798c630013c0

package main

import (
    "fmt"
    "time"
)

func UnluckyDays(year int) int {
	// Implement me! :)
  count:=0
  for i:=1; i<=12; i++{
    date:=time.Date(year,time.Month(i),13,0,0,0,0,time.UTC)
    if date.Weekday()==time.Friday{
      count++
    }
  }
  return count
  
}

func main(){
	n:=2015
	fmt.Println(UnluckyDays(n))
}


