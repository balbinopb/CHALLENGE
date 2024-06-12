
//https://www.codewars.com/kata/578aa45ee9fd15ff4600090d

package main

import "fmt"

func SortArray(array []int) []int {
	for i := 0; i < len(array); i++ {
		minIndex := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] && (array[minIndex]%2 == 1 && array[j]%2 == 1) {
				minIndex = j
			}
		}
		temp := array[i]
		array[i] = array[minIndex]
		array[minIndex] = temp
	}
	return array
}
func main() {
	arr := []int{5, 8, 6, 3, 4}
	fmt.Println(SortArray(arr))
}
