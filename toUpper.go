
//https://www.codewars.com/kata/57a0556c7cb1f31ab3000ad7

package main

import (
	"fmt"
	"strings"
	)

func MakeUpperCase(str string) string {
    return strings.ToUpper(str)
}

func main(){

	fmt.Println(MakeUpperCase("hElLo WoRlD"))

}