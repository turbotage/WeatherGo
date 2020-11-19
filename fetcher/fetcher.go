package fetcher

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
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

func serialReadLine() {

}

func rainFetch() {

}

func windFetch() {

}

// Wind Direction, Wind Speed, Gust
func fetchWind(serial *serial.Port, db *sql.DB) {

	buf := make([]byte, 128)

	n, err := serial.Write([]byte("1"))
	check(err)
	n, err = serial.Read(buf)
	check(err)
	fmt.Println(n)
	fmt.Println("%q", buf[:n])

}

func bme280Fetch() {

}

func fetchCycle(fI FetchingInfo) {
	done := false
	for i := 1; done; i++ {
		if (i % fI.rainupdatetime) == 0 {
			rainFetch()
		}
		if (i % fI.windupdatetime) == 0 {
			windFetch()
			gustFetch()
		}
		if (i % fI.bme280updatetime) == 0 {
			bme280Update()
		}
		time.Sleep(time.Second)
	}
}

/* "BeginFetching the function used to begin fetching" */
func BeginFetching(fetchingInfo FetchingInfo) {

	c := &serial.Config{Name: fetchingInfo.serialname, Baud: fetchingInfo.baud}
	serialport, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	/*
		dbname := "tcp(127.0.0.1:3306)/weather"
		db, err := sql.Open("mysql", "turbotage:klassuger@"+dbname)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		// Check that the database can be reached
		err = db.Ping()
		if err != nil {
			fmt.Println(err)
		}
	*/
	time.Sleep(5 * time.Millisecond)

}
