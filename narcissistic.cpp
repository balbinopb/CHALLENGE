
//https://www.codewars.com/kata/5287e858c6b5a9678200083c

#include<iostream>
#include<cmath>

using namespace std;

bool narcissistic( int value ){
  //Code away
  int n=value,count=0,sum=0;
  while(n>0){
    count++;
    n/=10;
  }
  int m=value;
 for(int i=0; i<count; i++){
   int temp=m%10;
   sum+=pow(temp,count);
   m/=10;
 }
  
  return sum==value;
  
}

int main(){
    int value = 153;
    cout<<narcissistic(value)<<endl;
    return 0;

}