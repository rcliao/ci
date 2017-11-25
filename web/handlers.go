package web

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
