package main

import "fmt"

func GetGrade(a,b,c int) rune {
    // Code here
  var grade rune
  avg:=(a+b+c)/3
  if 90<=avg && avg<=100{
    grade='A'
  }else if 80<=avg && avg<90{
    grade = 'B'
  }else if 70<=avg && avg<80{
    grade = 'C'
  }else if 60<=avg && avg<70{
    grade = 'D'
  }else if 0<=avg && avg<60{
    grade = 'F'
  }
  return grade
}

func main(){
	fmt.Println(GetGrade(80,46,90))
}