#include<iostream>
#include <cmath>
using namespace std;

double EasyLogs(double x, double a, double b) {
    double loga = log(a) / log(x); // Calculate log base x of a
    double logb = log(b) / log(x); // Calculate log base x of b
    double result = loga + logb;    // Add the logarithms
    return result;                  // Return the result
}
int main (){
    double x,a,b,result ;
    cin>>x>>a>>b;
    result=EasyLogs(x,a,b);
    cout<<result<<endl;
}