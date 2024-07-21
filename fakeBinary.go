
//https://www.codewars.com/kata/57eae65a4321032ce000002d

package main

import "fmt"

func FakeBin(x string) string {
      // your code here
    result := make([]byte, len(x))
    for i := 0; i < len(x); i++ {
        digit := int(x[i] - '0')
        if digit >= 5 {
            result[i] = '1'
        } else {
            result[i] = '0'
        }
    }
    return string(result)
}

func main(){
	fmt.Println(FakeBin("45385593107843568"))
}