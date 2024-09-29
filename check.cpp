
//https://www.codewars.com/kata/57cc975ed542d3148f00015b

#include <iostream>
#include <vector>
#include <string>

bool check(const std::vector<std::string>& seq, const std::string& elem) {
  for (int i=0; i<seq.size(); i++){
    if (seq[i]==elem){
      return true;
    }
  }
    return false;
}

int main(){
	std::vector<std::string> myVec={"a","b","c"};
	std::string elemt="b";
	std::cout<<check(myVec,elemt)<<std::endl;
}
