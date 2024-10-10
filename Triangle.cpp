
//https://www.codewars.com/kata/5a023c426975981341000014

#include<iostream>

using namespace std;

class Triangle {
public:
  static int otherAngle(int a, int b){
    return 180-(a+b);
  }
    
};

int main(){
	Triangle triangle;	
	cout<<triangle.otherAngle(30,60)<<endl;
}
