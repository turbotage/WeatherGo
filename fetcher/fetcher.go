package fetcher

import (
	"encoding/binary"
	"fmt"
	"math"
	"sync"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

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

func serialReadLine(s *serial.Port, db *sql.DB) {

}

func rainFetch(s *serial.Port, db *sql.DB) {

}

// Wind Direction, Wind Speed, Gust
func fetchWind(s *serial.Port, db *sql.DB) {

	buf := make([]byte, 128)

	n, err := s.Write([]byte("1"))
	check(err)
	n, err = s.Read(buf)
	check(err)
	fmt.Println(n)
	fmt.Println("%q", buf[:n])

}

func bme280Fetch(s *serial.Port, db *sql.DB) {

}

func fetchCycle(s *serial.Port, db *sql.DB) {
	done := false
	for i := 1; done; i++ {
		if (i % 1800) == 0 {
			rainFetch(s, db)
		}
		if (i % 600) == 0 {
			fetchWind(s, db)
		}
		if (i % 600) == 0 {
			bme280Fetch(s, db)
		}
		time.Sleep(time.Second)
	}
}

/* "BeginFetching the function used to begin fetching" */
func BeginFetching(wg *sync.WaitGroup, password string, serialname string, baud int) {
	defer wg.Done()

	c := &serial.Config{Name: serialname, Baud: baud}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("mysql", "turbotage:"+password+"@"+"tcp(127.0.0.1:3306)/weather")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Check that the database can be reached
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(5 * time.Millisecond)

	time.Sleep(2 * time.Second)

	fetchWind(s, db)

}
