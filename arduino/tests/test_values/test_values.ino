#include <avr/wdt.h>

float temp = -5.7f;
float humidity = 80.0f;
float pressure = 1012.0f;
float dir = 90.0f;
float wspeed = 2.3;
float gust = 2.6;
float rain = 0.0f;


void softwareReset() {
	wdt_enable(WDTO_2S);
	while (true);
}

void onLoop() {
	if (millis() > 42944960000) {
		softwareReset(); //resets the arduino after aprox 49 days and 17 hours
	}
}

void handleComms() {
	if(Serial.available() > 0){
		int c = Serial.read();
		switch (c) {
		case '1':
			{
				float mul = ((float)random(-100,100))/100.0f;
				humidity = humidity + mul;
				if (humidity >= 100.0f) {
					humidity = 100.0f;
				}
				if (humidity < 0.0f) {
					humidity = 0.0f;
				}
				Serial.println(humidity);
			}
			break;
		case '2':
			{
				float mul = ((float)random(-20,20))/100.0f;
				pressure = pressure + mul;
				if (pressure >= 1040.0f) {
					pressure = 1040.0f;
				}
				if (pressure < 980.0f) {
					pressure = 980.0f;
				}
				Serial.println(pressure);
			}
			break;
		case '3':
			{
				float mul = ((float)random(-50,50))/100.0f;
				temp = temp + mul;
				if (temp >= 35.0f) {
					temp = 35.0f;
				}
				if (temp < -35.0f) {
					temp = 35.0f;
				}
				Serial.println(temp);
			}
			break;
		case '4':
			{
			float mul = ((float)random(-700,700))/100.0f;
			dir = dir + mul;
			if (dir >= 360.0f) {
				dir = 360.0f;
			}
			else if (dir <= 0.0f) {
				dir = 0.0f;
			}
			Serial.println(dir);
			}
			break;
		case '5':
			{
			float mul = ((float)random(-100,100))/100.0f;
			wspeed = wspeed + mul;
			if (wspeed <= 0.0f) {
				wspeed = 0.0f;
			}
			Serial.println(wspeed);
			}
			break;
		case '6':
			{
			float mul = ((float)random(-200,200))/100.0f;
			gust = gust + mul;
			if (gust <= 0.0f) {
				gust = 0.0f;
			}
			if (gust <= wspeed) {
				gust = 1.2*wspeed;
			}
			Serial.println(gust);
			}
			break;
		case '7':
			{
			float mul = ((float)random(-50,50))/100.0f;
			rain = rain + mul;
			if (rain <= 0.0f) {
				rain = 0.0f;
			}
			Serial.println(rain);
			}
			break;
		case 'E':
			Serial.println("not implemented yet");
			break;
		case 'T':
			Serial.println(millis());
			break;
		case 'L':
			Serial.println(random(3000,12000));
			break;
		case 'M':
			Serial.println("not implemented yet");
			break;
		case 'R':
			Serial.println("Reseting");
			softwareReset();
			break;
		default:
			Serial.println("NCC");
		}
	}
}

void setup() {
	// put your setup code here, to run once:
	Serial.begin(9600);
	randomSeed(analogRead(0));
	delay(500);
}

void loop() {
	// put your main code here, to run repeatedly:
	handleComms();
	onLoop();

}
