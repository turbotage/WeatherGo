#include <math.h>

unsigned long startTime = 0;
unsigned long endTime = 0;

float doSomeWork(unsigned int workLoad) {
	float c = 2;
	for (int i = 0; i < 2; ++i){
		for(int j = 0; j < workLoad; ++j) {
			c += (workLoad % j) + sin(c) + 3.141592;
		}
	}
	return c;
}

void setup() {
	// put your setup code here, to run once:
	Serial.begin(9600);
}


void loop() {
	if (Serial.available() > 0) {
		int c = Serial.read();
		if (c == '1') {
			startTime = micros();
			// STARTS
			float a = doSomeWork(100);
			{
				a += doSomeWork(101);
				{
					a += doSomeWork(99);
					{
						a += doSomeWork(103);
					}
				}
			}
			// ENDS
			endTime = micros();
			Serial.println(endTime - startTime);
			Serial.println(a);
		}
	}
}
