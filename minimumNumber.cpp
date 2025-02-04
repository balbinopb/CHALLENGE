#include <bits/stdc++.h>

using namespace std; 

bool isPrime(int n){
  if(n<2) return false;
  if (n==2) return true;
  if (n%2==0) return false;
  
  for (int i=3; i<=sqrt(n); i+=2){
    if(n%i==0) return false;
  }
  
  return true;
}

int minimumNumber (vector <int> numbers )
{
  // Your Coed is here .. Enjoy !!!
  int sum=0;
  for(int i=0; i<numbers.size(); i++){
    sum+=numbers[i];
  }
  
  int needed = 0;
  while (!isPrime(sum + needed)) {
      needed++;
  }

  return needed;
}

int main(){
    vector<int> num = {3,1,2,3,1,1,1};
    cout<<minimumNumber(num)<<endl;
    return 0;

}