
//https://www.codewars.com/kata/57cc981a58da9e302a000214

#include <iostream>
#include <vector>

bool small_enough(std::vector<int> arr, int limit) {
  // your code here
  int sum=0;
  for (int i=0; i<arr.size(); i++){
    if (arr[i]>limit){
      return false;
      break;
    }
  }
  return true;
}

int main(){
	std::vector<int> vec={66,101};
	int limit=200;
	std::cout<<small_enough(vec,limit)<<std::endl;
}
