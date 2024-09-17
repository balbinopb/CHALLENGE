#include <iostream>
#include <vector>
#include <algorithm>

long sumTwoSmallestNumbers(std::vector<int> numbers)
{
    std::sort(numbers.begin(), numbers.end());
    return numbers[0] + numbers[1];
}

int main()
{
    std::vector<int> myVec = {19, 5, 42, 2, 77};
    std::cout << sumTwoSmallestNumbers(myVec) << std::endl;

    return 0;
}

