#include <iostream>
#include <vector>

std::vector<int> get_even_numbers(const std::vector<int>& arr) {
    std::vector<int> evens;
    for (size_t i = 0; i < arr.size(); ++i) {
        if (arr[i] % 2 == 0) {
            evens.push_back(arr[i]);
        }
    }
    return evens;
}

int main() {
    
    std::vector<int> numbers = {2, 4, 5, 6};

    std::vector<int> evens = get_even_numbers(numbers);

    std::cout << "Even numbers: ";
    for (int num : evens) {
        std::cout << num << " ";
    }
    std::cout << std::endl;

    return 0;
}

