//https://www.codewars.com/kata/57cfdf34902f6ba3d300001e

package main

import (
	"fmt"
	"sort"
	"strings"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)
	str := arr[0]
	var result strings.Builder
	for i, char := range str {
		if i > 0 {
			result.WriteString("***")
		}
		result.WriteRune(char)
	}
	return result.String()
}

func main(){
	arr:= []string{ "bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps" }
	fmt.Println(TwoSort(arr))
}