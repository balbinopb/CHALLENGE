
//https://www.codewars.com/kata/5ad0d8356165e63c140014d4

#include <iostream>
using namespace std;

int final_grade(int exam, int projects) {
    int total;
    if (exam > 90 || projects > 10) {
        total = 100;
    } else if (exam > 75 && projects >= 5) {
        total = 90;
    } else if (exam > 50 && projects >= 2) {
        total = 75;
    } else {
        total = 0;
    }
    return total;
}

int main() {

    int grade = final_grade(100, 12);

    cout <<  grade << endl;

    return 0;
}

