#include<iostream>
using namespace std;
long long rowSumOddNumbers(unsigned n){
  //your code here
 return n * n * n;
}
int main(){
    int n ,result;
    cin>>n;
    result=rowSumOddNumbers(n);
    cout<<result;
}