#include<iostream>
using namespace std;

bool IsLeapYear(int year) {
  if (year % 4 == 0 && (year % 100 != 0 || year % 400 == 0)){
    return true;
  }else {
    return false;
  }
}
int main (){
    int year,result ;
    cin>>year;
    result=IsLeapYear(year);
    cout<<result<<endl;
}