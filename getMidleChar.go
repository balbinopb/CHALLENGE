
//https://www.codewars.com/kata/56747fd5cb988479af000028

package main

import "fmt"

func GetMiddle(s string) string {
    lengthS := len(s)
    if lengthS == 0 {
        return ""
    }
    middleIndex := lengthS / 2
    if lengthS%2 == 0 {
        return s[middleIndex-1 : middleIndex+1]
    } else {
        return string(s[middleIndex])
    }
}

func main(){
	str:="middle"
	fmt.Println(GetMiddle(str))
}