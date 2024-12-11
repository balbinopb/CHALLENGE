
//https://www.codewars.com/kata/5cd4aec6abc7260028dcd942

#include <iostream>

using namespace std;

namespace Solution {
    int ShortestStepsToNum(int num) {
        int count = 0;

        while (num > 1) {
            if (num % 2 != 0) {
                num--; 
            } else {
                num /= 2; 
            }
            count++;
        }

        return count;
    }
}

int main(){
    int solution = Solution::ShortestStepsToNum(16);
    cout << solution << endl;
    
}