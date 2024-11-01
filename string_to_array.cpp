
//https://www.codewars.com/kata/57e76bc428d6fbc2d500036d

#include <iostream>
#include <vector>
#include <string>
#include <sstream>

std::vector<std::string> string_to_array(const std::string& s) {
  std::vector<std::string> myVector;
  std::istringstream iss(s);
  std::string word;

  if (s.empty()) {
    return {""};
  }

  while (iss >> word) {
    myVector.push_back(word);
  }
  return myVector;
}

int main() {
  std::vector<std::string> myVec = string_to_array("Robin Singh");

  for (const auto& word : myVec) {
    std::cout << word << " ";
  }
  
  return 0;
}


