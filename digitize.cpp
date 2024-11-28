
//https://www.codewars.com/kata/5417423f9e2e6c2f040002ae

#include <iostream>
#include <vector>

using namespace std;

std::vector<int> digitize(int n) {
  // your code here
  std::vector<int> a;
  if (n==0){
    a.insert(a.begin(),n);
    return a;
  }
  
  while(n>0){
    int t=n%10;
    a.insert(a.begin(),t);
    n /=10;
    
  }
  return a;
}

int main(){
	
	vector<int> n=digitize(123);
	
	for (int i=0; i<n.size(); i++){
		cout<<n[i]<<" ";
	}
	
}
