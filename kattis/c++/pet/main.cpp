#include <iostream>

int main() {
    int g = 0;
    int r;

    for(int i=0;i<5;i++) {
        int s = 0;
        for (int j=0;j<4;j++) {
            int x;
            std::cin >> x;
            s+=x;
        }
        if (s > g) {
            g=s;
            r=i+1;
        }
    }
    std::cout << r << " " << g;
    return 0;
}