#include <iostream>
#include <fstream>
#include <vector>


void day_2(std::vector<std::string> vec){

  int horizontal = 0, depth = 0, aim = 0;

  for (std::string s : vec){
    int movement = stoi(s.substr(s.find(" ")));
    s = s.substr(0, s.find(" "));
          
    if (s == "forward"){
      horizontal += movement;
      depth += aim * movement;
    }

    if (s == "down"){
      aim += movement;
    }

    if (s == "up"){
      aim -= movement;
    }
  }

  std::cout << "Result 2: "<< horizontal * depth << std::endl;
}


void day_1(std::vector<std::string> vec){
  
  int horizontal = 0;
  int depth = 0;
  for (std::string s : vec){
    int movement = stoi(s.substr(s.find(" ")));
    s = s.substr(0, s.find(" "));
    
    if (s == "forward")
      horizontal += movement; 

    if (s == "down")
      depth += movement;

    if (s == "up")
      depth -= movement;  
  }
  
  std::cout << "Result: " << horizontal * depth << std::endl;
}


int main(){
  std::ifstream file("input.txt");
  std::string line;
  std::vector<std::string> submarine;

  
  while(std::getline(file, line)){
    submarine.push_back(line); 
  }
  
  day_1(submarine);
  day_2(submarine);

  file.close();
}
