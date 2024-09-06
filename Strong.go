
//https://www.codewars.com/kata/5a4d303f880385399b000001

package main

import "fmt"

func Strong(n int) string {
    // your code here
  saveN:=n
  count:=0
  for n!=0{
    temp :=n%10
    count +=factorial(temp)
    n /=10
  }
  if count==saveN{
    return "STRONG!!!!"
  }
  return "Not Strong !!"
}


func factorial(n int)int{
  if n==0 || n==1{
      return 1
  }
  return n*factorial(n-1)

}

func main(){

fmt.Println(Strong(145))


}