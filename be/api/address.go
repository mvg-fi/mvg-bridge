package api

import (
	"encoding/json"
	"net/http"

	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/store"
)

// Handle Address creation
// 1. Create deposit address
// 2. Return deposit trace id
func (a *API) AddressHandler(s *store.BadgerStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST supported", http.StatusMethodNotAllowed)
			return
		}

		var d constants.DepositReq
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		traceID, err := s.WriteDeposit(&constants.Deposit{
			Type:        d.Type,
			ToAddress:   d.ToAddress,
			FromAssetID: d.FromAssetID,
			ToAssetID:   d.ToAssetID,
			Memo:        d.Memo,
			Amount:      d.Amount,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(constants.DepositResp{
			TraceID: traceID,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}
