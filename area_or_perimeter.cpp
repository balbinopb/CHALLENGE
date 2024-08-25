#include<iostream>

using namespace std;


int area_or_perimeter(int l , int w) {
  return (l==w)? l*l : 2*(l+w) ;
}


int main(){
	
	cout<<area_or_perimeter(5,7);
}
