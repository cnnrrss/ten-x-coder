#include <algorithm>
#include <iostream>
#include <iomanip>
#include <vector>

using namespace std;

typedef vector<double> vd;

int main() {
    int x;
    cin >> x;
    // decrement test cases
    while (x--) {
        // Number of Students
        int n;
        cin >> n;
        // create vector
        vd v(n);
        // total grade
        double s = 0;
        // loop through number of students
        for (int i = 0; i < n; i++) {
            // read student grade
            cin >> v[i];
            s += v[i];
        }
        sort(v.begin(), v.end());
        // x = number of students > s / n
        double x = v.end() - upper_bound(v.begin(), v.end(), s / n);
        cout << fixed << setprecision(3) << x * 100 / n << "%" << endl;
    }
}