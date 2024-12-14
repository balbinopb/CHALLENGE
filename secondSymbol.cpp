
//https://www.codewars.com/kata/63f96036b15a210058300ca9

#include <iostream>
#include <string>

using namespace std;

int secondSymbol(const std::string& str, char symbol) {
  //your code here
  int sec=0,idx;
  for(int i=0; i<str.length(); i++){
    if(str[i]==symbol){
      idx=i;
      sec++;
    }
    if(sec==2){
      return idx;
    }
  }
  return -1;
}

int main(){
    string str = "hello world";
    char symbol = 'o';
    cout<<secondSymbol(str,symbol)<<endl;
    return 0;
}