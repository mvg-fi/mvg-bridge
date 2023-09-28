package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mvg-fi/mvg-bridge/providers"
)

// Get all supported asset
func (a *API) AssetsHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET supported", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		log.Printf("/assets => %+v\n", r)

		resp := providers.ReadAssets()

		json.NewEncoder(w).Encode(resp)
	})
}
