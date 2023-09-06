package api

import (
	"encoding/json"
	"net/http"

	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/providers"
)

func (a *API) PriceSimpleHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST supported", http.StatusMethodNotAllowed)
			return
		}

		var d constants.PriceSimpleReq
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncode(w).Encode()
		providers.GetPrice()
		// return status
		w.Write([]byte("xd"))
	})
}

func (a *API) PriceAllHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST supported", http.StatusMethodNotAllowed)
			return
		}

		var d constants.DepositResp
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			return
		}
		// Get status by trace ID
		// d.TraceID

		// return status
		w.Write([]byte("xd"))
	})
}
