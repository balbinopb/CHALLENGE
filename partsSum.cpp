
//

#include<iostream>
#include<vector>

using namespace std;

std::vector<unsigned long long> partsSum(const std::vector<unsigned long long>& ls){
  //your code here
  vector<unsigned long long> sums(ls.size()+1,0);
  for (int i=ls.size()-1; i>=0; --i){
    sums[i]=sums[i+1]+ls[i];
  }  
  return sums;
}

int main(){
    vector<unsigned long long> ls = {4,6,3};
    vector<unsigned long long> res = partsSum(ls);
    for (auto i:res) cout<<i<<" ";
    return 0;
	
}