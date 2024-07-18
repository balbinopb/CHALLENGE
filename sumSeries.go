
//https://www.codewars.com/kata/555eded1ad94b00403000071

package main

import "fmt"

func SeriesSum(n int) string {
  // your code here
	start := 1
	difference := 3
	var sum float64

	for i := 0; i < n; i++ {
		denominator := start + i*difference
		term := 1.0 / float64(denominator)
		sum += term
	}
  return fmt.Sprintf("%.2f", sum)
}

func main(){
	fmt.Println(SeriesSum(5))
}