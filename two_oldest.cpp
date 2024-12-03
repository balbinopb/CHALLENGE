
//https://www.codewars.com/kata/511f11d355fe575d2c000001

#include <iostream>
#include <vector>
#include <array>
#include <algorithm>

using namespace std;

std::array<int, 2> two_oldest_ages(std::vector<int> ages)
{
    // your solution goes here
  sort(ages.begin(),ages.end());
  return {ages[ages.size()-2],ages[ages.size()-1]};
}

int main(){
  vector<int> ages = {1,5,87,45,8,8};
  auto result = two_oldest_ages(ages);
  cout<<result[0]<<" "<<result[1]<<endl;
}