
//

#include<iostream>

using namespace std;

int maxMultiple(int divisor, int bound) 
{
  // Your Code is Here ... Enjoy !!! 
  int temp;
  for (int i=divisor; i<=bound; i +=divisor){
      temp=i;
    if (i>temp){
      temp=i;
    }
  }
  return temp;
}


int main(){
	
	
	cout<<maxMultiple(7,17);
	
}
