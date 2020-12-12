
#include "def.h"

#define GUST_CONSTANT 1000 * 0.66698368f / GUST_TIME

class Anemometer {
private:
    //wind speed
    unsigned long m_TimeAtWindUpdate = 0;

    //wind gust
    float m_HighestGust = 0;
    bool m_GustHasUpdated = false;
    unsigned long m_TimeAtLastGustUpdate = 0;
    int m_GustPeriodCounter = 0;
	unsigned int m_GustRevs[NUM_GUST_COUNTERS];

public:

    void setup() {
		for (int i = 0; i < NUM_GUST_COUNTERS; ++i) {
			m_GustRevs[i] = 0;
		}
		m_TimeAtWindUpdate = millis();
		m_TimeAtLastGustUpdate = millis();
    }

    //===============================================================
    //Calculate the max wind speed/gust
    //===============================================================
    void updateWindGust(unsigned int revsSinceLastUpdate){
		for (int i = 0; i < NUM_GUST_COUNTERS; ++i) {
			m_GustRevs[i] += revsSinceLastUpdate;
		}
        unsigned long timer = (unsigned long)(millis() - m_TimeAtLastGustUpdate);
        int timu = timer % (int)(GUST_TIME / NUM_GUST_COUNTERS);
        if ((timu >= 0) && (timu <= 20)) {
            if(!m_GustHasUpdated){
                m_GustHasUpdated = true;
                float currentGust = GUST_CONSTANT * m_GustRevs[m_GustPeriodCounter];
                if(currentGust > m_HighestGust) {
                    m_HighestGust = currentGust;
                }
                m_GustRevs[m_GustPeriodCounter] = 0;
                ++m_GustPeriodCounter;
                if(m_GustPeriodCounter == NUM_GUST_COUNTERS){
                    m_GustPeriodCounter = 0;
                }
                m_TimeAtLastGustUpdate = millis();
            }
        }
        else {
            m_GustHasUpdated = false;
        }
    }

    //=======================================================
    // Calculate the wind speed, and display it (or log it, whatever).
    // 1 rev/sec = 1.492 mph = 2.40114125 kph, 0.6669836806 m/s
    //=======================================================
    float getWindSpeed(unsigned int windSpeedRevs) {
		unsigned long currTime = millis();
	    float windspeed = (1000.0f*0.66698368f * windSpeedRevs) / (currTime - m_TimeAtWindUpdate));
	    m_TimeAtWindUpdate = currTime; //set last update to now so next  wind callculation will be since this one
	    return windspeed;
    }

    //=======================================================
    // Get the highest gust speed during this cycle
    //=======================================================
    float getGust() {
        float tempGust = m_HighestGust;
        m_HighestGust = 0;
        return tempGust;
    }

};





