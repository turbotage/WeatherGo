#include <Adafruit_BME280.h>

#define BME_UPDATE_TIME 10000

//Pins
#define BME_CS 10
#define BME_MOSI 11
#define BME_MISO 12
#define BME_SCK 13

#define SEALEVELPRESSURE_HPA (1013.25)

Adafruit_BME280 bme(BME_CS,BME_MOSI,BME_MISO,BME_SCK);

void setup() {
	// put your setup code here, to run once:
	Serial.begin(9600);
	if(!bme.begin()) {
		Serial.println("BME not working!");
	}

}

void loop() {
  // put your main code here, to run repeatedly:
  	if (Serial.available() > 0) {
	  	int c = Serial.read();
	  	switch (c) {
		  	case '1':
				Serial.println(bme.readTemperature());
		  		break;
			case '2':
				Serial.println(bme.readPressure() /1000.0f);
				break;
			case '3':
				Serial.println(bme.readHumidity());
				break;
		
	  	}
  	}
}
