
//https://www.codewars.com/kata/551f37452ff852b7bd000139

#include<iostream>
#include <cstdint>
#include <string>

std::string add_binary(uint64_t a, uint64_t b) {
    uint64_t sum=a+b;
    std::string binary;
    do {
      binary=((sum%2==0)? "0" : "1") +binary;
      sum /=2;
    }while(sum>0);
    return binary;
    
}

int main(){
	
	std::cout<<add_binary(1,1)<<std::endl;
}
