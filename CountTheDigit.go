
//https://www.codewars.com/kata/566fc12495810954b1000030

package main

import "fmt"

import "strconv" // to convert 

func NbDig(n int, d int) int {
    sumD := 0
    dStr := strconv.Itoa(d)  // Convert the digit to a string for comparison

    for i := 0; i <= n; i++ {
        k := i * i
        kStr := strconv.Itoa(k)  // Convert the square to a string
        for _, char := range kStr {
            if string(char) == dStr {
                sumD++
            }
        }
    }
    return sumD
}

func main (){
	fmt.Println(NbDig(10,1))
}