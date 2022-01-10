#include <iostream>
#include <fstream>
#include <stdexcept>
#include <string>
#include <vector>

std::vector<int> depthIncreases(std::ifstream & file){
  int count = 0;
  std::string depth;
  std::vector<int> totalDepths;
  while(std::getline(file, depth)){
    totalDepths.push_back(stoi(depth)); 
  }


  for (int i = 0; i < totalDepths.size()-1; i++){
    if (totalDepths.at(i) < totalDepths.at(i+1)){
      count++;
    }
  }
  std::cout << "Increases: " << count << std::endl;

  return totalDepths;
}


void threeMeasure(std::vector<int> vec){
  
  int count = 0; 
  std::vector<int> measures = vec;
 
  for (int i = 0; i < vec.size()-(vec.size()%3)-1; i++){
    int paneOne = vec.at(i) + vec.at(i+1) + vec.at(i+2);
    int paneTwo = vec.at(i+1) + vec.at(i+2) + vec.at(i+3);
    if (paneOne < paneTwo){
      count++;
    }
    std::cout << vec.at(i) << "\n";
  }

  std::cout << "Three measure increases: " << count << std::endl;
}


int main(){
  std::ifstream submarine("input.txt");

  if (submarine.fail()) {
      std::cerr << "File cannot be opened for reading." << std::endl;
      exit(1);
  }
  //checking if ssh key broke
  threeMeasure(depthIncreases(submarine));

  submarine.close(); 
}
