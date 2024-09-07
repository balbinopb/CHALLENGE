
//https://www.codewars.com/kata/58b8c94b7df3f116eb00005b

#include <iostream>
#include <string>
#include <vector>

std::string reverse_letter(const std::string &str)
{
    std::string reverse;
    for (int i = str.length() - 1; i >= 0; --i) {
        if (str[i] >= 'a' && str[i] <= 'z')  {
            reverse.push_back(str[i]);
        }
    }
    return reverse;
}


int main(){
	
	std::cout<<reverse_letter("ab23c")<<std::endl;
	std::cout<<reverse_letter("ultr53o?n")<<std::endl;
	
}
