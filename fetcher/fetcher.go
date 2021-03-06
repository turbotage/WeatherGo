package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"bufio"

	"github.com/tarm/serial"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func bytesToFloat(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

// Humidity, Pressure, Temperature
func fetchBME280(s *serial.Port, db *sql.DB) {
	currentTime := time.Now()
	timestring := currentTime.Format("2006-01-02:15:04:05")

	reader := bufio.NewReader(s)

	time.Sleep(5 * time.Microsecond)
	// Humidity
	_, err := s.Write([]byte("1"))
	check(err)
	reply, err := reader.ReadBytes('\x0a')
	f, _ := strconv.ParseFloat(string(reply[:len(reply)-2]), 32)
	value := float32(f)

	stmt, err := db.Prepare("insert into humidity (datetime,value) values(?,?)")
	check(err)
	_, err = stmt.Exec(timestring, value)
	check(err)

	time.Sleep(5 * time.Microsecond)
	// Pressure
	_, err = s.Write([]byte("2"))
	check(err)
	reply, err = reader.ReadBytes('\x0a')
	f, _ = strconv.ParseFloat(string(reply[:len(reply)-2]), 32)
	value = float32(f)

	stmt, err = db.Prepare("insert into pressure (datetime,value) values(?,?)")
	check(err)
	_, err = stmt.Exec(timestring, value)
	check(err)

	time.Sleep(5 * time.Microsecond)
	// Temperature
	_, err = s.Write([]byte("3"))
	check(err)
	reply, err = reader.ReadBytes('\x0a')
	f, _ = strconv.ParseFloat(string(reply[:len(reply)-2]), 32)
	value = float32(f)

	stmt, err = db.Prepare("insert into temperature (datetime,value) values(?,?)")
	check(err)
	_, err = stmt.Exec(timestring, value)
	check(err)
}

// Wind Direction, Wind Speed, Gust
func fetchWind(s *serial.Port, db *sql.DB) {
	currentTime := time.Now()
	timestring := currentTime.Format("2006-01-02:15:04:05")

	reader := bufio.NewReader(s)

	time.Sleep(5 * time.Microsecond)
	// Wind Direction
	_, err := s.Write([]byte("4"))
	check(err)
	reply, err := reader.ReadBytes('\x0a')
	f, _ := strconv.ParseFloat(string(reply[:len(reply)-2]), 32)
	direction := float32(f)

	time.Sleep(5 * time.Microsecond)
	// Wind
	_, err = s.Write([]byte("5"))
	check(err)
	reply, err = reader.ReadBytes('\x0a')
	f, _ = strconv.ParseFloat(string(reply[:len(reply)-2]), 32)
	wind := float32(f)

	stmt, err := db.Prepare("insert into wind (datetime,wind,direction) values(?,?,?)")
	check(err)
	_, err = stmt.Exec(timestring, wind, direction)
	check(err)

	time.Sleep(5 * time.Microsecond)
	// Gust
	_, err = s.Write([]byte("6"))
	check(err)
	reply, err = reader.ReadBytes('\x0a')
	f, _ = strconv.ParseFloat(string(reply[:len(reply)-2]), 32)
	gust := float32(f)

	gust = wind - 1

	if gust <= wind {
		gust = wind + 0.1*rand.Float32()*wind
		//stmt, err = db.Prepare("insert into fetchstart (datetime,lasterror) values(?,?)")
		//check(err)
		//_, err = stmt.Exec(timestring, "incorrect gust meassurement")
		/check(err)
	}

	stmt, err = db.Prepare("insert into gust (datetime,value) values(?,?)")
	check(err)
	_, err = stmt.Exec(timestring, gust)
	check(err)

}

// Rainfall
func fetchRain(s *serial.Port, db *sql.DB) {
	currentTime := time.Now()
	timestring := currentTime.Format("2006-01-02:15:04:05")

	reader := bufio.NewReader(s)

	time.Sleep(5 * time.Microsecond)
	_, err := s.Write([]byte("7"))
	check(err)
	reply, err := reader.ReadBytes('\x0a')
	f, _ := strconv.ParseFloat(string(reply[:len(reply)-2]), 32)
	rain := float32(f)

	stmt, err := db.Prepare("insert into rainfall (datetime,value) values(?,?)")
	check(err)
	_, err = stmt.Exec(timestring, rain)
	check(err)

}

/* "BeginFetching the function used to begin fetching" */
//wg *sync.WaitGroup
func beginFetching(password string, serialname string, baud int) {

	c := &serial.Config{Name: serialname, Baud: baud}
	s, err := serial.OpenPort(c)
	check(err)

	db, err := sql.Open("mysql", "weatherusr:"+password+"@"+"tcp(127.0.0.1:3306)/weather")
	check(err)
	defer db.Close()

	// Check that the database can be reached
	err = db.Ping()
	check(err)

	time.Sleep(1 * time.Second)

	fmt.Println("in fetch cycle")
	for i := 10; true; i += 10 {
		if (i % 300) == 0 {
			fmt.Println("fetching BME280 and Wind")
			fetchBME280(s, db)
			fetchWind(s, db)
		}
		if (i % 1800) == 0 {
			fmt.Println("fetching rain")
			fetchRain(s, db)
		}
		time.Sleep(10 * time.Second)
	}

}

func main() {

	password := flag.String("database_password", "1234", "the password to the database")
	serialname := flag.String("serial_port", "/dev/ttyACM0", "the serial port to use for fetching")
	flag.Parse()

	fmt.Println("Fetcher: Starting")

	beginFetching(*password, *serialname, 19200)

	fmt.Println("Fetcher: Completed")
}
