#define PIN_ANEMOMETER 2
#define PIN_RAINGAUGE 3

volatile unsigned long nextAnemometerTime = 0;
volatile int windSpeedRevs = 0;

volatile unsigned long nextRainGaugeTime = 0;
volatile int raingaugeCount = 0;

//=======================================================
// Interrupt handler for anemometer. Called each time the reed
// switch triggers (one revolution). won't work if windspeed exceed 166,7430555556 m/s
//=======================================================
void countAnemometer() {
    if(nextAnemometerTime == 0 || nextAnemometerTime < millis()){
        windSpeedRevs++;
        nextAnemometerTime = millis() + 4; //this means that the windstation won't give accurate meassurments if wind speeds exceed 50 m/s;
    }
}

//=======================================================
// Interrupt handler for rain gauge. Called each time the reed
// switch triggers (one drop).
//=======================================================
void countRainGauge() {
  if (nextRainGaugeTime == 0 || nextRainGaugeTime) {
    raingaugeCount++;
    nextRainGaugeTime = millis() + 20;
  }
}

void setup() {
  // put your setup code here, to run once:
  Serial.begin(9600);
  pinMode(PIN_ANEMOMETER, INPUT);
  pinMode(PIN_RAINGAUGE, INPUT);
  digitalWrite(PIN_ANEMOMETER, HIGH);
  digitalWrite(PIN_RAINGAUGE, HIGH);
  attachInterrupt(digitalPinToInterrupt(PIN_ANEMOMETER), countAnemometer, FALLING);
  attachInterrupt(digitalPinToInterrupt(PIN_RAINGAUGE), countRainGauge, FALLING);
}

void loop() {
  // put your main code here, to run repeatedly:
  if (Serial.available() > 0) {
    int c = Serial.read();
    switch(c) {
      case '1':
        Serial.println(windSpeedRevs);
        break;
      case '2':
        Serial.println(raingaugeCount);
        break;
    }
  }
}
