
//

#include <iostream>
#include <vector>
#include<algorithm>

using namespace std; 

int minimumSteps (vector <int> num ,  int input)
{
  // Your Code is Here ... Enjoy !!!
  
  sort(num.begin(),num.end());
  int res=num[0],count=0;  
  for (int i=1; i<num.size(); i++){
    res +=num[i];
    count++;
    if (res>=input) break;

  }
  
  return count ;
}

int main(){
  vector<int> num = {4,6,3};
  int input = 7;
  cout<<minimumSteps(num,input)<<endl;
  return 0;
}