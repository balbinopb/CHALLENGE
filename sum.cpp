#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int sum(vector<int> numbers)
{
    if (numbers.size() < 3) {
        return 0;
    }

    sort(numbers.begin(), numbers.end());

    int sum = 0;

    for (int i = 1; i < numbers.size() - 1; i++) {
        sum += numbers[i];
    }

    return sum;
}

int main(){
	vector <int> myVector={ 6, 2, 1, 8, 10 };
	cout<<sum(myVector)<<endl;
}
