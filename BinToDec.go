
//https://www.codewars.com/kata/57a5c31ce298a7e6b7000334

package main

import (
	"fmt"
    	"strconv"
)

func BinToDec(bin string) int {
    sumDec := 0
    length := len(bin)
    for i := 0; i < length; i++ {
        mood, err := strconv.Atoi(string(bin[length-1-i]))
        if err != nil {
            return 0
        }
        sumDec += mood << i 
    }
    return sumDec
}

func main(){

fmt.Println(BinToDec("101010"))

}