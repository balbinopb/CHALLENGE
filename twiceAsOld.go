
//https://www.codewars.com/kata/5b853229cfde412a470000d0/go

package main

import (
	"fmt"
	"math"
)

func TwiceAsOld(dadYearsOld, sonYearsOld int) int { 
  // Implement me
 return int(math.Abs(float64(dadYearsOld-sonYearsOld*2)))

  return 0;
}
func main(){
	fmt.Println(TwiceAsOld(42,21))
	fmt.Println(TwiceAsOld(22,21))
}