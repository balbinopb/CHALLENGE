

//https://www.codewars.com/kata/5672a98bdbdd995fad00000f

package main

import "fmt"

func rps(p1, p2 string) string {
	if p1 == "scissors" && p2 == "paper" {
		return "Player 1 won!"
	} else if p1 == "paper" && p2 == "scissors" {
		return "Player 2 won!"
	} else if p1 == "scissors" && p2 == "rock" {
		return "Player 2 won!"
	} else if p1 == "rock" && p2 == "scissors" {
		return "Player 1 won!"
	} else if p1 == "paper" && p2 == "rock" {
		return "Player 1 won!"
	} else if p1 == "rock" && p2 == "paper" {
		return "Player 2 won!"
	} else if p1 == p2 {
		return "Draw!"
	}
	return "Invalid input!"
}

func main() {
	fmt.Println(rps("rock", "scissors"))
}
