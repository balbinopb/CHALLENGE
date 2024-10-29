
//https://www.codewars.com/kata/5a4138acf28b82aa43000117
#include<iostream>
#include <vector>
#include <climits> // For INT_MIN

using namespace std; 

int adjacentElementsProduct(vector<int> inputArray) 
{
    int mxProduct = INT_MIN; 
    for (int i = 0; i < inputArray.size() - 1; i++) { 
        int product = inputArray[i] * inputArray[i + 1]; 
        if (product > mxProduct) {
            mxProduct = product;
        }
    }
    return mxProduct; 
}

int main(){
	
	std::vector<int> arr={1,2,3};
	std::cout<<adjacentElementsProduct(arr)<<std::endl;
}
