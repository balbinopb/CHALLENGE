
//https://www.codewars.com/kata/534ea96ebb17181947000ada

#include<iostream>

using namespace std;

int break_chocolate(int n, int m){
  if (n*m >0){
    return n*m-1;
  }
  return 0;
}

 int main(){
 	
 	cout<<break_chocolate(7,4)<<endl;
 	
 }
