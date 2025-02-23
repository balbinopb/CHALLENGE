#include<iostream>

using namespace std;

int getRealFloor(int f) {
    if(f<=0){
      return f;
    }
    if(f<13){
      return f-1;
    }
    
   
    return f-2;
  }

  int main(){
    cout<<getRealFloor(5)<<endl;
    cout<<getRealFloor(45)<<endl;
  }