#include <vector>
#include <algorithm>
int sumOfDifferences(const std::vector<int>& arr){
  //your code here
  if(arr.empty()) return 0;
  int mx = *std::max_element(arr.begin(), arr.end());
  int mn = *std::min_element(arr.begin(), arr.end());
  
  return mx-mn;
}

int main(){
    std::vector<int> arr = {1, 2, 10};
    sumOfDifferences(arr);
    return 0;
}