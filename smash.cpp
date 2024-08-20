
//https://www.codewars.com/kata/53dc23c68a0c93699800041d
#include<iostream>
#include <vector>
#include <string>

using namespace std;

string smash(const vector<string>& words) {
    string concatenate;
    for (size_t i = 0; i < words.size(); ++i) {
        concatenate += words[i];
        if (i != words.size() - 1) {
            concatenate += " ";
        }
    }
    return concatenate;
}

int main(){
	vector<string> strArr = {"hello", "world", "this", "is", "great"};

	cout<<smash(strArr)<<endl;
}
