package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	socketio "github.com/googollee/go-socket.io"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

/* BeginServer the web server*/
func BeginServer(wg *sync.WaitGroup, password string) {
	defer wg.Done()

	server, err := socketio.NewServer(nil)
	check(err)

	db, err := sql.Open("mysql", "weatherusr:"+password+"@"+"tcp(127.0.0.1:3306)/weather")
	check(err)
	defer db.Close()

	rows, err := db.Query("select * from wind")
	check(err)

	columns, err := rows.Columns()
	check(err)
	fmt.Println(columns)
	fmt.Println(columns[0])

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected", s.ID())
		return nil
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("/asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
