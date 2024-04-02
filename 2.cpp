#include<iostream>
#include <cmath>
using namespace std;

int multiply(int number) {
    int count = 0;
    int num = number;

    if (num == 0) {
        count = 1;
    }

    while (num != 0) {
        num /= 10;
        count++;
    }

    int base = 5;
    int result = 1;
    for (int i = 0; i < count; i++) {
        result *= base;
    }

    return number * result;
}
int main (){
    double number,result ;
    cin>>number;
    result=multiply(number);
    cout<<result<<endl;
}