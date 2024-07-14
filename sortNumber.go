//https://www.codewars.com/kata/5174a4c0f2769dd8b1000003

package main

import (
	"fmt"
	"sort"
	)

func SortNumbers(numbers []int) []int {
  sort.Ints(numbers)
  return numbers// your code here
//   for i:=0; i<len(numbers); i++{
//     minIdx:=i
//     for j:=i+1; j<len(numbers); j++{
//       if numbers[i]>numbers[j]{
//         minIdx=j
//       }
//     }
//     numbers[i],numbers[minIdx]=numbers[minIdx],numbers[i]
//   }
}

func main(){
	arr:= []int{1, 2, 3, 10, 5}
	fmt.Println(SortNumbers(arr))
}