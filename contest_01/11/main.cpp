#include <iostream>
#include <math.h>
#include <string>
using namespace std;

int main() {
	int a, b, c;
	cin >> a >> b >> c;
	if (abs(a - b) < abs(a - c)) {
		cout << "B " + to_string(abs(a - b));
	}
	else {
		cout << "C " + to_string(abs(a - c));
	}

}