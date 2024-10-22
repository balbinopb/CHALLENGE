
//https://www.codewars.com/kata/55c28f7304e3eaebef0000da

#include<iostream>
#include <vector>

std::vector<int> createVector(const int n)
{
	std::vector<int> res;

	for (int i = 1; i <= n; i++)
	{
		res.push_back(i);
	}

	return res;
}

int main(){
	
	std::vector<int> Vector=createVector(5);
	for(int i=0; i<Vector.size(); i++){
		std::cout<<Vector[i]<<" ";
	}
	
	
}
