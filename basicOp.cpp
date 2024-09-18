
//https://www.codewars.com/kata/57356c55867b9b7a60000bd7

#include <iostream>

using namespace std;

int basicOp(char op, int val1, int val2) {
  // Your code here
  int result;

switch (op) {
  case '+':
    result=val1+val2;
    break;
  case '-':
    result=val1-val2;
    break;
  case '/':
    result = val1/val2;
    break;
  case '*':
    result = val1*val2;

}
  return result;


}

int main(){
	
	cout<<basicOp('+',4,7)<<endl;
	
}
