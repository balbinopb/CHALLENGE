#include <iostream>
#include <sstream>
#include <vector>

using namespace std;

string reverse_words(const string& str) {
    istringstream iss(str);
    vector<string> ws;
    string word;
    
    while (iss >> word) {
        ws.push_back(word);
    }

    string reversed;
    for (int i = ws.size() - 1; i >= 0; --i) {
        reversed += ws[i];
        if (i > 0) reversed += " ";
    }

    return reversed;
}

int main() {
    string str = "Hello World";
    cout << reverse_words(str) << endl;
    return 0;
}