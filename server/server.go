package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

/* BeginServer the web server*/
func beginServer(password string) {

	db, err := sql.Open("mysql", "weatherusr:"+password+"@"+"tcp(127.0.0.1:3306)/weather")
	check(err)
	defer db.Close()

	rows, err := db.Query("select * from wind")
	check(err)

	columns, err := rows.Columns()
	check(err)
	fmt.Println(columns)
	fmt.Println(columns[0])

	//http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("/asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))

}

func main() {
	var password = flag.String("database_password", "1234", "the password to the database")

	flag.Parse()

	fmt.Println("Server: Starting")

	beginServer(*password)

	fmt.Println("Server: Completed")
}
