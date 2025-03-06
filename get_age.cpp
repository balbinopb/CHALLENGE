#include <iostream>
#include <string>

using namespace std;

int get_age(const std::string& she_said) {
    return she_said[0]-'0';
}

int main() {
    return get_age("I am 20 years old");
}