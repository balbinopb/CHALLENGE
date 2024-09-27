
//https://www.codewars.com/kata/583710ccaa6717322c000105

#include<iostream>
using namespace std;

int simpleMultiplication(int a){
    //Your code
  return a%2==0? 8*a : 9*a;
}

int main(){
	
	cout<<simpleMultiplication(5)<<endl;
	
}
