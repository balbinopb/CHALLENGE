
//https://www.codewars.com/kata/54ba84be607a92aa900000f1

#include<iostream>

bool is_isogram(std::string str) {
  for (int i = 0; i < str.length(); i++)
	{
		for (int j = i + 1; j < str.length(); j++)
		{
			if (tolower(str[j]) == tolower(str[i]))
			{
				return false;
			}
		}
	}
	return true;
}

int main(){
	std::cout<<is_isogram("Dermatoglyphics")<<std::endl;
}
