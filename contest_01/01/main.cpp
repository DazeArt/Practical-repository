#include <iostream>
#include <math.h>
using namespace std;

int main() {
	float sql = sqrt(12.0);
	float s = 1.0 - (1.0 / 9.0) + (1.0 / 45.0) - (1.0 / (7.0*27.0)) + (1.0 / (3.0 * 3.0 * 3.0 * 3.0 * 3.0 * 3.0)) - (1.0 / (3.0 * 3.0 * 3.0 * 3.0 * 3.0 * 11.0));
	cout << sql * s;
}