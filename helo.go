
//https://www.codewars.com/kata/55a70521798b14d4750000a4

package main

import "fmt"

func Greet(name string) string {
	// fill in solution here
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}

func main() {
	fmt.Println(Greet("jay"))
}
