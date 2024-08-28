
//https://www.codewars.com/kata/57f609022f4d534f05000024

#include <iostream>
#include <vector>

int stray(std::vector<int> numbers) {
    for (int i = 0; i < numbers.size(); i++) {
        bool is_stray = true;
        for (int j = 0; j < numbers.size(); j++) {
            if (i != j && numbers[i] == numbers[j]) {
                is_stray = false;
                break;
            }
        }
        if (is_stray) {
            return numbers[i];
        }
    }
    return 0;
}


int main(){
	std::vector<int> arr = {1, 1, 2};
	std::cout<<stray(arr)<<std::endl;
}
