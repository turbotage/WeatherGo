
#include "def.h"

class RainGauge {
private:
    unsigned long m_TimeAtRainUpdate = 0;
public:

    void setup() {
      	m_TimeAtRainUpdate = millis();
    }

    //=======================================================
    // Calculate the rain , and display it (or log it, whatever).
    // 1 bucket = 0.2794 mm
    //=======================================================
    float getRainFall(unsigned int numRainDrops) {
        float x = 0.2794*numRainDrops;
        float timer = (millis() - m_TimeAtRainUpdate)/1000; //get seconds since last raincallculation
        m_TimeAtRainUpdate = millis(); //set last update to now so next rain callculation will be since this one
        return (3600 * x)/timer; //take volume per time and convert to right format (mm/h)
    }

};
