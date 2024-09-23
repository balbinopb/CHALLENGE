
//https://www.codewars.com/kata/5601409514fc93442500010b

#include <iostream>
#include <vector>

bool betterThanAverage(std::vector<int> classPoints, int yourPoints) {
  // your code here
  int avg=0;
  int len=classPoints.size();
  for (int i=0; i<len; i++){
    avg +=classPoints[i];
  }
  return yourPoints >= avg/(len-1);
}


int main(){
	std::vector<int> myVec={2,3};
	int n=5;
	std::cout<<betterThanAverage(myVec,n)<<std::endl;
	
}
