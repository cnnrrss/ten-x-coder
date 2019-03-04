#include <iostream>

int main() {
    int N;
    std::cin >> N;
    if (N < 10000000)
        std::cout << (N % 2 == 1 ? "Alice" : "Bob") << std::endl;
    return 0;
}
