#include <Adafruit_BME280.h>

#include "def.h"



class BME280 {
private:
    Adafruit_BME280 m_Bme;

    bool m_HasUpdated = false;
    
    //humidity
    int m_HumidityCounter = 0;
    float m_Humidity = 0;
    //temperature
    int m_TemperatureCounter = 0;
    float m_Temperature = 0;
    //pressure
    int m_PressureCounter = 0;
    float m_Pressure = 0;
public:

    BME280() 
    : m_Bme(BME_CS, BME_MOSI, BME_MISO, BME_SCK)
    {

    }

    bool setup() {
        if(!m_Bme.begin()) {
            return false;
        }
        return true;
    }

    void updateHumidityTempPressure() {
        unsigned long timu = millis() % BME_UPDATE_TIME;
        if ((timu >= 0) && (timu <= 40)) {
            if (!m_HasUpdated) {
                m_HasUpdated = true;
                m_Humidity += m_Bme.readHumidity();
                ++m_HumidityCounter;
                m_Pressure += (m_Bme.readPressure()/1000.0f);
                ++m_PressureCounter;
                m_Temperature += m_Bme.readTemperature();
                ++m_TemperatureCounter;
            }
        }
        else {
            m_HasUpdated = false;
        }
    }

    float getHumidity() {
        if(m_HumidityCounter != 0){
            float humid = m_Humidity / m_HumidityCounter;
            m_Humidity = 0;
            m_HumidityCounter = 0;
            return humid;
        }
        else {
            return m_Bme.readHumidity();
        }
    }

    float getPressure() {
        if(m_PressureCounter != 0) {
            float press = m_Pressure / m_PressureCounter;
            m_Pressure = 0;
            m_PressureCounter = 0;
            return press;
        }
        else {
            return (m_Bme.readPressure() / 1000.0f);
        }
    }

    float getTemperature() {
        if(m_TemperatureCounter != 0) {
            float temp = m_Temperature / m_TemperatureCounter;
            m_Pressure = 0;
            m_PressureCounter = 0;
            return temp;
        }
        else {
            return m_Bme.readTemperature();
        }
    }

};
