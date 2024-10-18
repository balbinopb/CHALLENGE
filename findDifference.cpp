
//https://www.codewars.com/kata/58cb43f4256836ed95000f97

#include <iostream>
#include <array>

int findDifference(std::array<int, 3> a, std::array<int, 3> b) {
  // your code here
  if (a[0]*a[1]*a[2]<b[0]*b[1]*b[2]){
    return -1*(a[0]*a[1]*a[2]-b[0]*b[1]*b[2]);
  }
  return a[0]*a[1]*a[2]-b[0]*b[1]*b[2];
}

int main(){
	
	std::array<int,3> arr1={3, 2, 5};
	std::array<int,3> arr2={1, 4, 4};
	std::cout<<findDifference(arr1,arr2)<<std::endl;
	
}
