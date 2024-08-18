package main

import "fmt"

func countDigit(n int)int{
	countdgt:=0
	for n!=0{
		n /=10
		countdgt++
	}
	return countdgt
}


func main(){

	dgt:=2
	fmt.Println(countDigit(dgt))
}