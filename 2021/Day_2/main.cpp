#include <iostream>
#include <fstream>
#include <vector>





int main(){
  std::ifstream file("input.txt");
  std::string line;
  std::vector<std::string> submarine;

  
  while(std::getline(file, line)){
    submarine.push_back(line); 
  }
  
  int horizonal = 0;
  int depth = 0;
  for (std::string s : submarine){
    int movement = stoi(s.substr(s.find(" ")));
    s = s.substr(0, s.find(" "));
    
    if (s == "forward")
      horizonal += movement; 

    if (s == "down")
      depth += movement;

    if (s == "up")
      depth -= movement;  
  }
  
  std::cout << "Result: " << horizonal * depth << std::endl;

}
