
//https://www.codewars.com/kata/54da5a58ea159efa38000836

package main

import "fmt"

func FindOdd(seq []int) int {
// your code here
	result := 0
	for i := 0; i < len(seq); i++{
		result ^= seq[i]
	}
	return result
}

func main(){
arr:= []int{1,2,2,3,3,3,4,3,3,3,2,2,1}

fmt.Println(FindOdd(arr))

}