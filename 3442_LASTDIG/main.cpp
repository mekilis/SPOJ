#include <iostream>
#include <cmath>
#include <string>
#include <sstream>

using namespace std;

int main() {
	
	int powerMatrix[21][5];
	
	for (int i = 0; i <= 20; i++) {
		for (int j = 0; j <= 4; j++) {
			int power = pow(i, j);
			ostringstream o;
			o << power;
			string powerStr = o.str();
			int lastDigit = (int)(powerStr.length() - 1);
			powerMatrix[i][j] = (powerStr[lastDigit]) - '0';;
		}
	}
	
	int t;
	cin >> t;
	
	for (int tt = 0; tt < t; tt++) {
		int a, b;
		cin >> a;
		cin >> b;
		
		int key = b;
		if (key != 0) {
			key = b % 4;
			if (key == 0) {
				key = 4;
			}
		}
		
		cout << powerMatrix[a][key] << endl;
	}
	
	return 0;
}
