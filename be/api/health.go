package api

import (
	"encoding/json"
	"net/http"

	"github.com/mvg-fi/mvg-bridge/constants"
)

// TODO
// Return overall bridge health
func (a *API) HealthHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET supported", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		resp := &constants.Health{}

		json.NewEncoder(w).Encode(resp)
	})
}
