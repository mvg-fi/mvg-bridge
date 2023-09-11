package api

import (
	"encoding/json"
	"fmt"
	"log"
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

		w.Header().Set("Content-Type", "application/json")

		var req constants.PriceSimpleReq
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)

			return
		}
		log.Printf("/price/simple => %+v\n", req)

		amount, fee := providers.GetPriceSimple(req.FromAssetID, req.ToAssetID, req.Amount, req.Except, req.Cex)

		json.NewEncoder(w).Encode(&constants.PriceSimpleResp{
			Amount: fmt.Sprintf("%v", amount),
			Fee:    fmt.Sprintf("%v", fee),
		})
	})
}

func (a *API) PriceAllHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST supported", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		var req constants.PriceAllReq
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Wrong input", http.StatusBadRequest)
			return
		}
		log.Printf("/price/all => %+v\n", req)

		prices := providers.GetPriceAll(req.FromAssetID, req.ToAssetID, req.Amount, req.Except, req.Cex)

		json.NewEncoder(w).Encode(prices)
	})
}
