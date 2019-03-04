#include <iostream>
#include <stdio.h>


using namespace std;

int main() {
    int x, y;
    scanf("%d %d", &x, &y);

    char c[1000];
    for(int i=1;i<=x;i++){
        (void)scanf("%s", c);
    }

    printf("%d\n", y);
}
