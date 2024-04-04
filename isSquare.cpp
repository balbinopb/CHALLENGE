#include <iostream>
#include <cmath>

using namespace std;

bool is_square(int n)
{
  // TODO
  int root ;
  root= sqrt(n);
    return root*root==n;
}
int main(){
    int n;
    cin>>n;
    cout<<is_square(n);

}