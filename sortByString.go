
//https://www.codewars.com/kata/57ea5b0b75ae11d1e800006c

package main

import "fmt"

func SortByLength(arr []string) []string {
  n := len(arr)
  for i := 0; i < n; i++ {
      minIdx := i
      for j := i + 1; j < n; j++ {
          if len(arr[j]) < len(arr[minIdx]) {
              minIdx = j
          }
      }
      arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
  return arr
}

func main(){
	str:= []string{"Telescopes", "Glasses", "Eyes", "Monocles"}
	fmt.Println(SortByLength(str))
}