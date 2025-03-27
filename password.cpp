

//day 161
#include <bits/stdc++.h>

using namespace std;

bool password(const std::string &pwd)
{
    if (pwd.length() < 8) return false;

    bool hasUpper = false, hasLower = false, hasDigit = false;

    for (char i : pwd) {
        if (std::isupper(i)) hasUpper = true;
        if (std::islower(i)) hasLower = true;
        if (std::isdigit(i)) hasDigit = true;
    }

    return hasUpper && hasLower && hasDigit;
}

int main(){
    string password1="Password123";
    if(password(password1)){
        cout<<"Valid password"<<endl;
    }else{
        cout<<"Invalid password"<<endl;
    }
}