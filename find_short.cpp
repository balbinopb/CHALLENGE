#include <iostream>
#include <string>

int find_short(std::string str)
{
    int minLength = 1000;  
    int currentLength = 0;

    for (char c : str) {
        if (c == ' ') {
            if (currentLength > 0 && currentLength < minLength) {
                minLength = currentLength;
            }
            currentLength = 0;
        } else {
            currentLength++;
        }
    }

    // Handle the last word
    if (currentLength > 0 && currentLength < minLength) {
        minLength = currentLength;
    }

    return minLength;
}

int main() {
    std::string str = "bitcoin take over the world maybe who knows perhaps";
    std::cout << find_short(str) << std::endl;
    return 0;
}

