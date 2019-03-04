#include <iostream>

int main() {
    int n, m;
    std::cin >> n >> m;
    if (n<=1000 && m<=1000 && n!=m) {
        if (m >= n) {
            if (m - n == 1) std::cout << "Dr. Chaz will have 1 piece of chicken left over!\n";
            else std::cout << "Dr. Chaz will have " << m - n << " pieces of chicken left over!" << std::endl;
        } else {
            if (n - m == 1) std::cout << "Dr. Chaz needs 1 more piece of chicken!\n";
            else std::cout << "Dr. Chaz needs "  << n - m << " more pieces of chicken!" << std::endl;
        }
    }

    return 0;
}
