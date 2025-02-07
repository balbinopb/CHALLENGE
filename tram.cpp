#include <iostream>
#include <vector>

using namespace std;

int tram(int stops, const std::vector<int>& a, const std::vector<int>& b) {
  // "GoodLuck";
  int mx=0,tmp=0;
  for(int i=0; i<stops; i++){
    mx -=a[i];
    mx +=b[i];
    if(mx>tmp){
        tmp=mx;
    }
  }
  return tmp;
  
}

int main() {
    //example
    vector<int> a = {0, 2, 3, 4, 5};
    vector<int> b = {2, 3, 4, 5, 6};
    cout << tram(5, a, b) << endl;  // 6
  
}