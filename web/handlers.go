package web

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rcliao/e2etest/github"
)

// Health checks health of the service
func Health(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping; err != nil {
			http.Error(
				w,
				"Unhealthy",
				http.StatusInternalServerError,
			)
			return
		}
		fmt.Fprintln(w, "Healthy")
	})
}

// Authorize redirects the user to Github authorization page
func Authorize(api *github.API) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, api.AuthorizationLink(), 302)
	})
}

// GetToken uses Github API to get access token and store token into DB
func GetToken(api *github.API) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		log.Println("code", code)
		log.Println("token", api.GetToken(code))
	})
}

// Hook handles the webhook from Github API call
func Hook() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var event github.Event
		err := decoder.Decode(&event)
		if err != nil {
			http.Error(
				w,
				"failed to parse body as JSON",
				http.StatusBadRequest,
			)
			return
		}

		fmt.Println("Got webhook", event)
	})
}
