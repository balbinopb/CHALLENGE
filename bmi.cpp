
//https://www.codewars.com/kata/57a429e253ba3381850000fb

#include <iostream>
#include <string>

std::string bmi(double w, double h) 
{
  float bmi=w/(h*h);
  
  if (bmi <= 18.5) return "Underweight";

  if (bmi <= 25.0) return "Normal";

  if (bmi <= 30.0) return "Overweight";

  //if (bmi > 30) return "Obese"


  
    return "Obese";
}

int main(){
	
	double w=4,h=9;
	std::cout<<bmi(w,h)<<std::endl;
	
}
