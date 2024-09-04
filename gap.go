
//https://www.codewars.com/kata/561e9c843a2ef5a40c0000a4

package main

import "fmt"

func Gap(g, m, n int) []int {
    var prevPrime int
    for i := m; i <= n; i++ {
        if isPrime(i) {
            if prevPrime != 0 && i-prevPrime == g {
                return []int{prevPrime, i}
            }
            prevPrime = i
        }
    }
    return nil
}

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    if n == 2 {
        return true
    }
    if n%2 == 0 {
        return false
    }
    for i := 3; i*i <= n; i += 2 {
        if n%i == 0 {
            return false
        }
    }
    return true
}


func main(){

fmt.Println(Gap(2,100,110))

}