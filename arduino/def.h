
//============================= WIND VANE =======================
#define VANE_UPDATE_TIME 5000 //meassurements every VANE_UPDATE_TIME for wind directions
#define PIN_DIRECTION A0 //Analog 0  

//============================= ANEMOMETER ======================
#define GUST_TIME 2000
#define NUM_GUST_COUNTERS 32
#define PIN_ANEMOMETER 2  //Digital 2


//============================= RAINGAUGE =======================
#define PIN_RAINGAUGE 3 //Digital 3


//============================ BME 280 ==========================
#define BME_UPDATE_TIME 20000 //meassurements every BME_UPDATE_TIME for humidity, pressure and temperature
#define SEALEVELPRESSURE_HPA (1013.25)
//Pins
#define BME_CS 10
#define BME_MOSI 11
#define BME_MISO 12
#define BME_SCK 13