
//https://www.codewars.com/kata/56f69d9f9400f508fb000ba7

#include <iostream>
#include <vector>

std::vector<int> MonkeyCount(int n) {
  // your code here
  std::vector<int> vec={};
  
  for(int i=1; i<=n; i++){
    vec.push_back(i);
  }
  return vec;
}

int main(){
	std::vector <int> myVector=MonkeyCount(10);
	for (int i=0; i<myVector.size(); i++){
	    std::cout<<myVector[i]<<" ";
	}
	
	
}
