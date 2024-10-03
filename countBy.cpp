#include <iostream>
#include <vector>
std::vector<int> countBy(int x,int n){
  std::vector<int> myVec;
  for (int i=x; i<=x*n; i+=x){
    myVec.push_back(i);
  }
  return myVec;
}

int main(){
	
	std::vector<int> vec = countBy(5, 4);
	for (int i = 0; i < vec.size(); i++) {
        std::cout << vec[i] << std::endl;
    }
}
