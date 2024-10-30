
//https://www.codewars.com/kata/5a805d8cafa10f8b930005ba

package main

import (
	"fmt"
	"math"
)

func NearestSq(n int) int {
    root := int(math.Round(math.Sqrt(float64(n))))
    return root * root
}

func main() {
    fmt.Println(NearestSq(10))
}
