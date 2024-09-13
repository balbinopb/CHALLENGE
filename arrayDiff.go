
//https://www.codewars.com/kata/523f5d21c841566fde000009

package main

import "fmt"

func ArrayDiff(a, b []int) []int {
  // your code here
  
  // Create a map to store the elements of array b
	bMap := make(map[int]bool)
	for _, num := range b {
		bMap[num] = true
	}

	// Create a result slice for elements that are in a but not in b
	var result []int
	for _, num := range a {
		if !bMap[num] {
			result = append(result, num)
		}
	}

	return result
  
}

func main(){

a := []int{1,2,2,2,3}
b := []int{2}

fmt.Println(ArrayDiff(a,b))

}