
//https://www.codewars.com/kata/568dcc3c7f12767a62000038

#include <iostream>
using namespace std;


bool set_alarm(const bool& employed,const bool& vacation){

  return (employed==true && vacation==false)? true :false;
}

int main(){
	cout<<set_alarm(true,true)<<endl;
	cout<<set_alarm(true,false)<<endl;
}
