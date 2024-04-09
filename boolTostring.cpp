#include<iostream>
#include <string>

using namespace std;
string boolean_to_string(bool b){
  //Your code here
  if (b) {
        return "true";
    } else {
        return "false";
    }
}
int main(){
    bool n;
    cin>>n;
    string result=boolean_to_string(n);
    cout<<result;

}