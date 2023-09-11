package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/store"
	"github.com/mvg-fi/mvg-bridge/users"
)

func (a *API) OrderHandler(s *store.BadgerStore) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST supported", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		var req constants.OrderReq
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		log.Printf("/order/new => %+v\n", req)

	})
}

func handleCreateOrder(ctx context.Context, p *users.Proxy, s *store.BadgerStore, req *constants.OrderReq) {
	// Create order
	// Get random user
	p.GetADeposit(ctx, s, req.FromAssetID)
}
