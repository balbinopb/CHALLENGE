
//https://www.codewars.com/kata/55f9bca8ecaa9eac7100004a

#include<iostream>

using namespace std;

int past(int h, int m, int s) {
  // your code here
  return (3600*h+60*m+s)*1000;
}

int main(){
	cout<<past(1,1,1)<<endl;
	
}
