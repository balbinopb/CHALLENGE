#include<iostream>

using namespace std;

string greet(string name) 
{
  //debug
  //return "Hello, " + name + "!";
  if(name == "Johnny") {
    return "Hello, my love!";
  }
  return "Hello, " +name +"!";
}

int main(){
	cout<<greet("Beatriz")<<endl;
	cout<<greet("Johnny")<<endl;
	
}
