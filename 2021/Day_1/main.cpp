#include <iostream>
#include <fstream>
#include <string>
#include <vector>

void depthIncreases(std::ifstream & file){
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
}


void threeMeasure(std::string file){

}


int main(){
  std::ifstream submarine("input.txt");

  if (submarine.fail()) {
      std::cerr << "File cannot be opened for reading." << std::endl;
      exit(1);
  }

  depthIncreases(submarine);

  submarine.close(); 
}
