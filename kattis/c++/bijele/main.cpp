#include <iostream>
#include <list>
#include <stdio.h>

int main() {
    int s[] = {1, 1, 2, 2, 2, 8};
    for(int i=0;i<6;i++){
        int x;
        std::cin >> x;
        if(s[i] == x)
            std::cout << "0";
        else if(s[i] > x)
            std::cout << s[i] - x;
        else
            std::cout << (x - s[i]) * -1;
        if(i < 6) std::cout << " ";
    }
    std::cout << std::endl;
}