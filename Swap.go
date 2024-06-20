
//https://www.codewars.com/kata/5388f0e00b24c5635e000fc6

package main

import (
    "fmt"
)

func swapValues(values [2]int) [2]int {
    temp := values[0]
    values[0] = values[1]
    values[1] = temp
    return values
}

func main() {
    values := [2]int{3, 2}
    swappedValues := swapValues(values)
    fmt.Println(swappedValues)
}
