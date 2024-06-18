
//https://www.codewars.com/kata/5526fc09a1bbd946250002dc

package main

import "fmt"

func FindOutlier(integers []int) int {
	var singleOdd, singleEven int
	even := 0
	odd := 0

	for i := 0; i < len(integers); i++ {
		if integers[i]%2 == 0 {
			even++
			singleEven = integers[i]
		} else {
			odd++
			singleOdd = integers[i]
		}

		if even > 1 && odd == 1 {
			return singleOdd
		}
		if odd > 1 && even == 1 {
			return singleEven
		}
	}

	if even == 1 {
		return singleEven
	}
	return singleOdd
}

func main() {
	arr := []int{2, 4, 0, 100, 4, 11, 2602, 36}
	fmt.Println(FindOutlier(arr))
}
