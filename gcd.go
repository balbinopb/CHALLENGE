package main

import "fmt"

func main() {
    var A, B int
    fmt.Scan(&A, &B)
    fmt.Println(gcd(A, B))
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a % b)
}
