package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rcliao/e2etest/web"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := mux.NewRouter()
	db := getDB(os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))

	r.HandleFunc("/health", web.Health(db)).Methods("GET", "HEAD")

	log.Println("Running web server at port 8000")
	http.ListenAndServe(":8000", r)
}

func getDB(username, password, host string) *sql.DB {
	defaultProtocol := "tcp"
	defaultPort := "3306"

	sqlDSN := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/",
		username,
		password,
		defaultProtocol,
		host,
		defaultPort,
	)

	db, err := sql.Open("mysql", sqlDSN)
	if err != nil {
		panic(err)
	}

	return db
}
