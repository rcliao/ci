package web

import (
	"database/sql"
	"fmt"
	"net/http"
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
