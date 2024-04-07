#include <iostream>
#include <vector>

using namespace std;

int sum(vector<int> nums) {
  int sum=0;
  for (int i=0; i<nums.size(); i++){
    sum +=nums[i];
  }
  return sum;
  
  // your code here
}
int main (){
    vector<int> nums[];
    cin>>nums;
    cout<<sum(nums);
}
