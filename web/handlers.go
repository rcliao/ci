package web

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rcliao/e2etest"
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
func GetToken(api *github.API, tokenDao e2etest.TokenDAO) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		token := api.GetToken(code)
		if token == "" {
			http.Error(
				w,
				"failed to grab token",
				http.StatusInternalServerError,
			)
			return
		}
		err := tokenDao.StoreToken(token)
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to store token",
				http.StatusInternalServerError,
			)
			return
		}
		fmt.Fprintln(w, "Successfuly retrieve and store token. CI is ready.")
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

		log.Println("Got webhook", event)
	})
}
