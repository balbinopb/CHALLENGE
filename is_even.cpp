
#include <iostream> 
#include <cmath>  

bool is_even(double n) {
   
    return std::fmod(n, 2) == 0;
}

int main(){
	
	std::cout<<is_even(4.52)<<std::endl;
	
}
