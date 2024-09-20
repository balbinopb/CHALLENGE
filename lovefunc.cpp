
//https://www.codewars.com/kata/555086d53eac039a2a000083

#include<iostream>

using namespace std;

bool lovefunc(int f1, int f2) {
  return (f1%2==0 != f2%2==0) ;
}

int main(){
	
	cout<<lovefunc(1,4)<<endl;
	
}
