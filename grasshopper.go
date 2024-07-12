
//https://www.codewars.com/kata/55f73be6e12baaa5900000d4

package main

import "fmt"

func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
  // your code
  return laLigaGoals+copaDelReyGoals+championsLeagueGoals
}

func main(){
	fmt.Println(Goals(40,30,20))
}