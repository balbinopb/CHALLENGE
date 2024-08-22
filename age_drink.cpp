#include <iostream>
using namespace std;

string people_with_age_drink(int age) {
    if (age < 14) {
        return "drink toddy";
    } else if (age < 18) {
        return "drink coke";
    } else if (age < 21) {
        return "drink beer";
    } else {
        return "drink whisky";
    }
}

int main(){
	cout<<people_with_age_drink(37)<<endl;
}
