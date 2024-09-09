#include<iostream>
#include<string>

int square_digits(int num) {
   if (num < 0) {
        num = -num;
    }

    if (num == 0) {
        return 0;
    }

    std::string result;
    while (num > 0) {
        int digit = num % 10;
        num /= 10;
        int squared_digit = digit * digit;
        result = std::to_string(squared_digit) + result;
    }

    return std::stoi(result);
}

int main(){
	std::cout<<square_digits(9119 )<<std::endl;
	
}
