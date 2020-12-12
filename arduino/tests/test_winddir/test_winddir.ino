#define PIN_DIRECTION A0 
#include <math.h>

long readVcc() {
    long result;
    // Read 1.1V reference against AVcc
    ADMUX = _BV(REFS0) | _BV(MUX3) | _BV(MUX2) | _BV(MUX1);
    delay(2); // Wait for Vref to settle
    ADCSRA |= _BV(ADSC); // Convert
    while (bit_is_set(ADCSRA,ADSC));
    result = ADCL;
    result |= ADCH<<8;
	result = 1125300L / result; // Back-calculate AVcc in mV
	return result;
}

double voltageReading() {
	double vcc;
    unsigned int ADCValue;
    vcc = readVcc()/1000.0;
	ADCValue = analogRead(PIN_DIRECTION);
    return (ADCValue / 1024.0) * vcc;
}

float m_CorrespX[16] = { 
	1, 0.92387953, 0.70710678, 0.38268343,        // N, NNE, NE, ENE
	0, -0.38268343, -0.70710678, -0.92387953,     // E, ESE, SE, SSE
	-1, -0.92387953, -0.70710678, -0.38268343,    // S, SSW, SW, WSW
	0, 0.38268343, 0.70710678, 0.92387953         // W, WNW, NW, NNW
};

float m_CorrespY[16] = { 
    0, -0.38268343, -0.70710678, -0.92387953,     // N, NNE, NE, ENE
    -1, -0.92387953, -0.70710678, -0.38268343,     // E, ESE, SE, SSE
    0, 0.38268343, 0.70710678, 0.92387953,         // S, SSW, SW, WSW
    1, 0.92387953, 0.70710678, 0.38268343         // W, WNW, NW, NNW
};

String dir[16] = {"N","NNE","NE","ENE","E","ESE","SE","SSE","S","SSW","SW","WSW","W","WNW","NW","NNW"};
double volts[16] = {1.0408,0.7648,1.5152,1.3167,3.1272,2.9819,4.5939,4.0469,4.3277,3.4612,3.8485,2.0719,2.3284,0.5664,0.6099,0.4841};

void setup() {
  // put your setup code here, to run once:
  Serial.begin(9600);
}

int getDirectionIndex(double reading) {
	for (int i = 0; i < 16; ++i) {
		if ( (volts[i]*0.976 < reading) && (reading < volts[i]*1.0232) ) {
			return i;
		}
	}
	return -1;
}

int c = 0;
void loop() {
	if (Serial.available() > 0) {
		c = Serial.read();
	}
	delay(500);

	if (c == '1'){
		Serial.println(voltageReading(), 5);
	}
	else if (c == '2') {
		int i = getDirectionIndex(voltageReading());
		if (i == -1) {
			Serial.println("MISSED READING");
		}else {
			Serial.println(dir[i]);
		}
	}
	else if (c == '3') {
		int i = getDirectionIndex(voltageReading());
		if (i == -1) {
			Serial.println("MISSED READING");
		}
		else {
			double x = m_CorrespX[i];
			double y = m_CorrespY[i];
			float degree;
			if (y < 0.0) {
				degree = -1*180*atan2(y,x)/3.141592f;
			}
			else {
				degree = 360 - 180*atan2(y,x)/3.141592f;
			}
			Serial.println(degree);
			Serial.println(dir[i]);
		}
	}
}
