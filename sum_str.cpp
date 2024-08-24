#include <bits/stdc++.h>

using namespace std;

string sum_str(const string& a, const string& b) {
    int intA = a.empty() ? 0 : stoi(a);
    int intB = b.empty() ? 0 : stoi(b);
    int sum = intA + intB;
    return to_string(sum) ;
}

int main(){
	cout<<sum_str("4","5");
}
