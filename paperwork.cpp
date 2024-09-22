
//https://codewars.com/kata/55f9b48403f6b87a7c0000bd

#include <iostream>

using namespace std;


int paperwork(int n, int m){
  if (n<1 || m<1){
    return 0;
  }
  return n*m;
    
}

int main(){
	
	cout<<paperwork(5,5)<<endl;
	
}
