package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rcliao/e2etest/dao"
	"github.com/rcliao/e2etest/github"
	"github.com/rcliao/e2etest/web"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := mux.NewRouter()
	db := getDB(os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	api := github.NewAPI(
		os.Getenv("GITHUB_CLIENT_ID"),
		os.Getenv("GITHUB_CLIENT_SECRET"),
		os.Getenv("PUBLIC_URL"),
	)
	service := dao.New(db)

	r.HandleFunc("/health", web.Health(db)).Methods("GET", "HEAD")
	r.HandleFunc("/api/authorize", web.Authorize(api)).Methods("GET")
	r.HandleFunc("/api/github/callback", web.GetToken(api, service)).Methods("GET")
	r.HandleFunc("/api/webhook", web.Hook()).Methods("POST")

	log.Println("Running web server at port 8000")
	http.ListenAndServe(":8000", r)
}

func getDB(username, password, host, database string) *sql.DB {
	defaultProtocol := "tcp"
	defaultPort := "3306"

	sqlDSN := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s",
		username,
		password,
		defaultProtocol,
		host,
		defaultPort,
		database,
	)

	db, err := sql.Open("mysql", sqlDSN)
	if err != nil {
		panic(err)
	}

	return db
}
