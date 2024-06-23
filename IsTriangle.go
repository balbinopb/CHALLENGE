
//https://www.codewars.com/kata/56606694ec01347ce800001b

package main

import"fmt"

func IsTriangle(a, b, c int) bool {
    if a <= 0 || b <= 0 || c <= 0 {
        return false
    }
    if a+b > c && a+c > b && b+c > a {
        return true
    }
    return false
}

func main(){
	fmt.Println(IsTriangle(3,5,6))
	fmt.Println(IsTriangle(4,8,2))
}