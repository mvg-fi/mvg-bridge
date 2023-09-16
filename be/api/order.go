package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mvg-fi/common/time"
	"github.com/mvg-fi/common/uuid"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/store"
	"github.com/mvg-fi/mvg-bridge/users"
)

func (a *API) OrderHandler(ctx context.Context, p *users.Proxy, s *store.BadgerStore) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST supported", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		var req constants.OrderNewReq
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		log.Printf("/order/new => %+v\n", req)
		order := handleCreateOrder(ctx, p, s, &req)

		json.NewEncoder(w).Encode(&constants.OrderNewResp{
			TraceID: order.TraceID,
			Address: order.Address,
			Memo:    order.Memo,
			Expire:  order.Expire,
		})
	})
}

func handleCreateOrder(ctx context.Context, p *users.Proxy, s *store.BadgerStore, req *constants.OrderNewReq) *constants.Order {
	// Get random user
	orderID := uuid.NewV4()
	entires := p.GetADeposit(ctx, s, req.FromAssetID, orderID)

	// Create order
	expire := time.UTCNowAddMinutes(constants.ExpirePeriod)
	order := constants.Order{
		FromAssetID: req.FromAssetID,
		ToAssetID:   req.ToAssetID,
		Amount:      req.Amount,
		Except:      req.Except,
		Cex:         req.Cex,
		TraceID:     orderID,
		Address:     entires.Destination,
		Memo:        entires.Tag,
		Expire:      expire,
		Status:      constants.PrefixOrder,
	}

	// Store order
	s.WriteOrder(&order)
	return &order
}
