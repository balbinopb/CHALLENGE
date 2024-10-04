
//https://www.codewars.com/kata/563e320cee5dddcf77000158

#include <iostream>
#include <vector>
#include <cmath>

int get_average(std::vector <int> marks)
{
  // Your Code is Here ... Enjoy !!!
  int sumMarks=0;
  int i=0;
  while ( i<marks.size()){
    sumMarks  +=marks[i];
    i++;
  }
  return std::floor(sumMarks/i);
}

int main(){
	
	std::vector<int> myVector={1, 5, 87, 45, 8, 8};
	std::cout<<get_average(myVector)<<std::endl;
	
}
