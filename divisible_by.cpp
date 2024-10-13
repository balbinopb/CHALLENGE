
//https://www.codewars.com/kata/55edaba99da3a9c84000003b

#include <iostream>
#include <vector>

std::vector<int> divisible_by(std::vector<int> numbers, int divisor)
{
    std::vector<int> nums = {};
    for (int i = 0; i < numbers.size(); i++) {
        if (numbers[i] % divisor == 0) {
            nums.push_back(numbers[i]);
        }
    }
    return nums;
}

int main(){
	std::vector<int> myVector={1, 2, 3, 4, 5, 6};
	int n=2;
	std::vector<int> vec=divisible_by(myVector,n);
	for (int i=0; i<vec.size(); i++){
		std::cout<<vec[i]<<" ";
	}
	
}
