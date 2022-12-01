#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>


using namespace std;


int main(){
    ifstream file("input.txt");
    string line;

    vector<int> k;  
    int sum = 0;

    while(getline(file, line)){
        if (line != "") sum += stoi(line);
        if (line == ""){
            k.push_back(sum);
            sum = 0;
        }
    }

    sum = 0;
    sort(k.begin(), k.end(), greater<int>());
    
    for (auto i = 0; i < 3; ++i)
        sum += k[i];

    cout << sum << endl;


}