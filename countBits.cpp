
//https://www.codewars.com/kata/526571aae218b8ee490006f4/train/cpp

#include <iostream>
#include <list>

using namespace std;

unsigned int countBits(unsigned long long n) {
    unsigned int count = 0;  
    
    while (n > 0) {
        count += n % 2;  
        n = n / 2;       
    }
    
    return count;
}

int main(){
	
	cout<<countBits(1234)<<endl;
}
