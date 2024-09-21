
//https://www.codewars.com/kata/56676e8fabd2d1ff3000000c

#include <iostream>
#include <vector>
#include <string>

std::string findNeedle(const std::vector<std::string>& haystack)
{
    // You may use the function std::to_string to convert numbers to strings for easy concatenation. 
    // E.g., ("hi" + std::to_string(123)) ==> "hi123". Of course, that's only one of many approaches.
  for (int i=0; i<haystack.size(); i++){
    if (haystack[i]=="needle"){
      return "found the needle at position " + std::to_string(i);
    }
  }  
    return 0;
}

int main(){
	
	std::vector<std::string> Vec= {"hay", "junk", "hay", "hay", "moreJunk", "needle", "randomJunk"};
	std::cout<<findNeedle(Vec)<<std::endl;
	
}
