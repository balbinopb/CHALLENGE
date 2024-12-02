
//https://www.codewars.com/kata/5b752a42b11814b09c00005d

#include<iostream>

using namespace std;

std::pair<int, int> solve (int a, int b){
    //..
  if (a==0 || b==0){
    return {a,b};
  }else if(a>=2*b){
    return solve(a-2*b,b);
  }else if (b>=a*2){
    return solve(a,b-2*a);
  }else{
    return {a,b};
  }
}

int main(){
  int a = 6;
  int b = 19;
  auto result = solve(a,b);
  cout<<result.first<<" "<<result.second<<endl;
}