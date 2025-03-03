#include <iostream>
#include <string>

std::string double_char(const std::string& str)
{
  std::string res="";
  for(int i=0; i<str.length(); i++){
    res +=str[i];
    res +=str[i];
  }
    return res;
}

int main()
{
  //double_char("Hello");
  std::cout<<double_char("Hello")<<std::endl;
  return 0;
}