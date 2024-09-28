
//https://www.codewars.com/kata/5949481f86420f59480000e7

#include <iostream>
#include <string>
#include <vector>

std::string odd_or_even(const std::vector<int> &arr)
{ 
  int sum=0;
  for (int i=0; i<arr.size(); i++){
    sum +=arr[i];
  }
    return sum%2==0? "even" : "odd"; // your code here
}

int main(){
	std::vector<int> Myvec={0,1,4};
	std::cout<<odd_or_even(Myvec)<<std::endl;
}
