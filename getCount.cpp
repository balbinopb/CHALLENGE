
//https://www.codewars.com/kata/54ff3102c1bad923760001f3

#include <iostream>
#include <string>

using namespace std;

int getCount(const string& inputStr){
  int num_vowels = 0;
  //your code here
  for (int i=0; i<inputStr.length(); i++){
    if (inputStr[i]=='a' || inputStr[i]=='e' || inputStr[i]=='i' || inputStr[i]=='o' || inputStr[i]=='u'){
      num_vowels++;
    }
  }
  return num_vowels;
}

int main(){
	
	cout<<getCount("abracadabra")<<endl;
	
}
