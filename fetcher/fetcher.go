package fetcher

import (
	"encoding/binary"
	"fmt"
	"math"
	"sync"
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

func fetchRain(s *serial.Port, db *sql.DB) {
	currentTime := time.Now()
	timestring := "'" + currentTime.Format("2006-01-02:15:04:05") + "'"

	reader := bufio.NewReader(s)

	time.Sleep(5 * time.Microsecond)
	_, err := s.Write([]byte("7"))
	check(err)
	reply, err := reader.ReadBytes('\x0a')
	rain := string(reply)

	stmt, err := db.Prepare("insert into rainfall (datetime,value) values(?,?)")
	check(err)
	_, err = stmt.Exec(timestring, rain)
	check(err)

}

// Wind Direction, Wind Speed, Gust
func fetchWind(s *serial.Port, db *sql.DB) {
	currentTime := time.Now()
	timestring := "'" + currentTime.Format("2006-01-02:15:04:05") + "'"

	reader := bufio.NewReader(s)

	time.Sleep(5 * time.Microsecond)
	// Wind Direction
	_, err := s.Write([]byte("4"))
	check(err)
	reply, err := reader.ReadBytes('\x0a')
	wind := string(reply)

	time.Sleep(5 * time.Microsecond)
	// Wind
	_, err = s.Write([]byte("5"))
	check(err)
	reply, err = reader.ReadBytes('\x0a')
	direction := string(reply)

	stmt, err := db.Prepare("insert into wind (datetime,wind,direction) values(?,?,?)")
	check(err)
	_, err = stmt.Exec(timestring, wind, direction)
	check(err)

	time.Sleep(5 * time.Microsecond)
	// Gust
	_, err = s.Write([]byte("6"))
	check(err)
	reply, err = reader.ReadBytes('\x0a')
	gust := string(reply)

	stmt, err = db.Prepare("insert into gust (datetime,gust) values(?,?)")
	check(err)
	_, err = stmt.Exec(timestring, gust)
	check(err)

}

func fetchBME280(s *serial.Port, db *sql.DB) {
	currentTime := time.Now()
	timestring := "'" + currentTime.Format("2006-01-02:15:04:05") + "'"

	reader := bufio.NewReader(s)

	time.Sleep(5 * time.Microsecond)
	// Humidity
	_, err := s.Write([]byte("1"))
	check(err)
	reply, err := reader.ReadBytes('\x0a')
	value := string(reply)

	stmt, err := db.Prepare("insert into humidity (datetime,value) values(?,?)")
	check(err)
	_, err = stmt.Exec(timestring, value)
	check(err)

	time.Sleep(5 * time.Microsecond)
	// Pressure
	_, err = s.Write([]byte("2"))
	check(err)
	reply, err = reader.ReadBytes('\x0a')
	value = string(reply)

	stmt, err = db.Prepare("insert into pressure (datetime,value) values(?,?)")
	check(err)
	_, err = stmt.Exec(timestring, value)
	check(err)

	time.Sleep(5 * time.Microsecond)
	// Temperature
	_, err = s.Write([]byte("3"))
	check(err)
	reply, err = reader.ReadBytes('\x0a')
	value = string(reply)

	stmt, err = db.Prepare("insert into temperature (datetime,value) values(?,?)")
	check(err)
	_, err = stmt.Exec(timestring, value)
	check(err)
}

/* "BeginFetching the function used to begin fetching" */
//wg *sync.WaitGroup
func BeginFetching(doneFetching chan bool, wg *sync.WaitGroup, password string, serialname string, baud int) {
	//defer wg.Done()

	c := &serial.Config{Name: serialname, Baud: baud}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("mysql", "weatherusr:"+password+"@"+"tcp(127.0.0.1:3306)/weather")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Check that the database can be reached
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("in fetch cycle")
	done := false
	for i := 0; done; i += 10 {
		if (i % 60) == 0 {
			fmt.Println("fetching BME280 and Wind")
			fetchBME280(s, db)
			fetchWind(s, db)
		}
		if (i % 180) == 0 {
			fmt.Println("fetching rain")
			fetchRain(s, db)
		}
		time.Sleep(10 * time.Second)
	}

	doneFetching <- true

}
