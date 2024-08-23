
//https://www.codewars.com/kata/5865918c6b569962950002a1

#include <iostream>
#include <string>

using namespace std;

unsigned int strCount(const string& word, char letter){
	// your code here
  int suml=0;
  for (int i=0; i<word.length(); i++){
    if( word[i]==letter){
      suml++;
    }
  }
  return suml; 
}

int main(){
	cout<<strCount("Hello",'o')<<endl;
	cout<<strCount("Hello",'l');
}
