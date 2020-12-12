#include <math.h>

#include "def.h"

class WindVane {
private:
    float m_X = 0.0f;
    float m_Y = 0.0f;

    int m_NumOfUpdateCycles = 0;
    bool m_WindDirHasUpdated = false;
	
	//int m_Sorted[16] = {65,83,92,126,185,245,288,407,464,603,634,706,788,830,889,947};
    //float m_Perc[16] = {0.05,0.045,0.05,0.10,0.18,0.078,0.078,0.064,0.064,0.025,0.024,0.056,0.025,0.025,0.031,0.031};
    //int m_DirValues[16] = { 788, 407, 464, 83, 92, 65, 185, 126, 288, 245, 634, 603, 947, 830, 889, 706 };
    
	//String dir[16] = {"N","NNE","NE","ENE","E","ESE","SE","SSE","S","SSW","SW","WSW","W","WNW","NW","NNW"};
    float m_Voltages[16] = {1.0408,0.7648,1.5152,1.3167,3.1272,2.9819,4.5939,4.0469,4.3277,3.4612,3.8485,2.0719,2.3284,0.5664,0.6099,0.4841};
    

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

    int getDirectionIndex(double reading) {
		for (int i = 0; i < 16; ++i) {
			if ( (m_Voltages[i]*0.976 < reading) && (reading < m_Voltages[i]*1.0232) ) {
				return i;
			}
		}
		return -1;
    }
    
    
public:
    
    void updateWindDirection() {
        unsigned long timu = millis() % VANE_UPDATE_TIME;
        if ((timu >= 0) && (timu <= 40)) {
			if (!m_WindDirHasUpdated) {
				m_WindDirHasUpdated = true;
				double reading = voltageReading();
				int dirIndex = getDirectionIndex(reading);
				if(dirIndex == -1){
					return;
				}
				m_X += m_CorrespX[dirIndex];
				m_Y += m_CorrespY[dirIndex];
				++m_NumOfUpdateCycles;
			}
        }
        else {
            m_WindDirHasUpdated = false;
        }
    }

    float getWindDirection() {
        if (m_NumOfUpdateCycles == 0) {
        	updateWindDirection();
        }
        float degree;

        if (m_Y < 0.0) {
        	degree = -1*180*atan2(m_Y, m_X)/3.141592;
        }
        else {
          	degree = 360 - 180*atan2(m_Y, m_X)/3.141592;
        }

        /*
        if ((m_X >= 0) && (m_Y >= 0)) {
  		    degree = 90 - atan2(m_Y, m_X);
  	    } 
        else if (m_Y < 0) {
  		    degree = (-1 * atan2(m_Y, m_X)) + 90;
  	    }
        else if ((m_X < 0) && (m_Y >= 0)) {
  		    degree = 460 - atan2(m_Y, m_X);
  	    }
        */

        m_X = 0;
        m_Y = 0;
        m_NumOfUpdateCycles = 0;

        return degree;
    }

};
