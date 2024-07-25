//https://www.codewars.com/kata/545a4c5a61aa4c6916000755


package main

import "fmt"

func Gimme(array [3]int) int {
  if (array[0] > array[1] && array[0] < array[2]) || (array[0] < array[1] && array[0] > array[2]) {
    return 0
  } else if (array[1] > array[0] && array[1] < array[2]) || (array[1] < array[0] && array[1] > array[2]) {
    return 1
  } else {
    return 2
  }
}

func main(){
arr:= [3]int{2, 3, 1}
arr2:= [3]int{5, 10, 14}

fmt.Println(Gimme(arr))
fmt.Println(Gimme(arr2))


}