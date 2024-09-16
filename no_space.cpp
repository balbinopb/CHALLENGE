#include <iostream>
#include <string>

std::string no_space(const std::string& x)
{
  std::string result;
  
  for (char c : x) {
        if (c != ' ') {
            result += c;
        }
    }
  return result;
}

int main() {
	std::string str="8 j 8   mBliB8g  imjB8B8  jl  B";
	std::cout<<no_space(str)<<std::endl;
	
}
