
//https://www.codewars.com/kata/57eb8fcdf670e99d9b000272

package main

import"fmt"

func High(s string) string {
    // your code here
  highestScore := 0
	highestWord := ""
	word := ""
	currentScore := 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char != ' ' {
			word += string(char)
			currentScore += int(char - 'a' + 1)
		} else {
			if currentScore > highestScore {
				highestScore = currentScore
				highestWord = word
			}
			word = ""
			currentScore = 0
		}
	}
	if currentScore > highestScore {
		highestScore = currentScore
		highestWord = word
	}

	return highestWord
}

func main(){
	fmt.Println(High("aa b"))
}