#include <iostream>
using namespace std;

int main() {
    int T, NG, NM, x;
    int NGMax = 0;
    int NMMax = 0;
    cin >> T;

    for (int t = 0; t < T; t++) {
        cin.get(); //trash
        cin >> NG;
        cin >> NM;

        for (int ng = 0; ng < NG; ng++) {
            cin >> x;
            if (x > NGMax) {
                NGMax = x;
            }
        }

        for (int nm = 0; nm < NM; nm++) {
            cin >> x;
            if (x > NMMax) {
                NMMax = x;
            }
        }

        if (NMMax > NGMax) {
            cout << "MechaGodzilla" << endl;
        } else {
            cout << "Godzilla" << endl;
        }
    }

    return 0;
}
