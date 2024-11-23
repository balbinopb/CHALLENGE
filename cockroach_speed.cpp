
//https://www.codewars.com/kata/55fab1ffda3e2e44f00000c6

#include<iostream>
#include<cmath>

int cockroach_speed(double s)
{
  //s in km 100 cm=1m 1km=
  double speedInCm=(100*1000*s)/3600;
    return std::floor(speedInCm); //Good Luck!
}

int main(){
	std::cout<<cockroach_speed(1.08)<<std::endl;
}
