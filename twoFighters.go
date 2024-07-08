package main

import "fmt"

type Fighter struct {
    Name            string
    Health          int
    DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
  // Your code goes here. Have fun!
  for fighter1.Health>0 && fighter2.Health>0{
    if firstAttacker==fighter1.Name{
      // Fighter1 attacks first
      fighter2.Health -= fighter1.DamagePerAttack
      if fighter2.Health <= 0 {
        return fighter1.Name
      }
      // Now Fighter2 attacks
      fighter1.Health -= fighter2.DamagePerAttack
      if fighter1.Health <= 0 {
        return fighter2.Name
      }
    }else if firstAttacker==fighter2.Name{
      // Fighter2 attacks first
      fighter1.Health -= fighter2.DamagePerAttack
      if fighter1.Health <= 0 {
        return fighter2.Name
      }
      // Now fighter1 attacks 
      fighter2.Health -= fighter1.DamagePerAttack
      if fighter2.Health <= 0 {
        return fighter1.Name
      }
    }
  }
  return ""
}

func main(){
	var Fighter1 , Fighter2 Fighter
	Fighter1.Name,Fighter1.Health,Fighter1.DamagePerAttack="lew",10,2
	Fighter2.Name,Fighter2.Health,Fighter2.DamagePerAttack="harry",5,4
	firstAttacker:=Fighter1.Name
	fmt.Println(DeclareWinner(Fighter1,Fighter2,firstAttacker))
}