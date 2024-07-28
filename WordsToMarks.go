
//https://www.codewars.com/kata/59706036f6e5d1e22d000016

package main

import "fmt"

func WordsToMarks(s string) int {
  count := 0
  for _, i := range s {
     count += int(i) - 'a' + 1;
  }
  return count
}

func main(){

fmt.Println(WordsToMarks("love"))

}