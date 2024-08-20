
//https://www.codewars.com/kata/53369039d7ab3ac506000467
#include <iostream>
#include <string>

using namespace std;

string bool_to_word(bool value)
{
  if (value){
    return "Yes";
  }else {
    return "No";
  }
}

int main(){
	cout<<bool_to_word(true)<<endl;
	cout<<bool_to_word(false)<<endl;
}
