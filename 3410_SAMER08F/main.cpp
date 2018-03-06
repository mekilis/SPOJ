#include <iostream>
using namespace std;

int main() {
	
	// your code here
	int N = 0;
	
	cin >> N;
	
	while (N != 0) {
		cout << ((2*N) + 1) * ((N + 1) * N) / 6 << endl;
		cin >> N;
	}

	return 0;
}
