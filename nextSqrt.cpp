#include<iostream>
#include <cmath>
using namespace std;

int FindNextSquare(int sq) {
  int root = static_cast<int>(std::sqrt(static_cast<double>(sq)));
  if (root*root == sq) {
        int nextRoot = root + 1;
        return nextRoot * nextRoot;
    } else {
        return -1;
    }
}
int main (){
    int sq ;
    cin>>sq;
    cout<<FindNextSquare(sq)<<endl;
}

