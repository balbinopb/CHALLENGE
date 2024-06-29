
//https://www.codewars.com/kata/57a083a57cb1f31db7000028

package main

import "fmt"
import "math"

func PowersOfTwo(n int) []uint64 {
    // Create a slice to store the results
    result := make([]uint64, n+1)
    for i := 0; i <= n; i++ {
        result[i] = uint64(math.Pow(2, float64(i)))
    }
    return result
}

func main(){
	fmt.Println(PowersOfTwo(2))
}