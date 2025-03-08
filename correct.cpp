#include <iostream>
#include <string>

std::string correct(std::string str){
  //your code here
  std::string res="";
  for(char c : str){
    if((c-'0')==1){
      res +='I';
    }else if((c-'0')==0){
      res+='O';
    }else if((c-'0')==5){
      res+='S';
    }else{
      res+=c;
    }
  }
  return res;
}

int main(){
    std::string str = "1F-RUDYARD K1PL1NG";
    std::cout << correct(str) << std::endl;
    return 0;
}