
//https://www.codewars.com/kata/52c31f8e6605bcc646000082

package main

import "fmt"

func TwoSum(numbers []int, target int) [2]int {
  for i:=0; i<len(numbers); i++{
        for j:=i+1; j<len(numbers); j++{
            if numbers[i] + numbers[j]==target {
                return [2]int{i, j}
            }
        }
    }
    return [2]int{}
}

func main(){
	arr:= []int{1,2,3}
	target:=4

	fmt.Println(TwoSum(arr,target))
}