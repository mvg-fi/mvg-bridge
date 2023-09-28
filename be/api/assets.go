package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mvg-fi/mvg-bridge/constants"
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

		var req constants.PriceSimpleReq
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)

			return
		}
		log.Printf("/price/simple => %+v\n", req)

		resp := providers.GetPriceSimple(req.FromAssetID, req.ToAssetID, req.Amount, req.Except, req.Cex)

		json.NewEncoder(w).Encode(resp)
	})
}
