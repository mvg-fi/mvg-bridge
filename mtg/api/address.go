package api

import (
	"net/http"

	"golang.org/x/time/rate"
)

// Handle Address creation
func (a *API) AddressHandler(w http.ResponseWriter, req *http.Request) {
	// 0. Rate limit
	// 1. Create deposit address
	// 2. Return deposit

	limiter := rate.NewLimiter(3, 5) // 3 request per sec, maxium 5 per sec
	if !limiter.Allow() {
		http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		return
	}

	if req.Method != http.MethodPost {
		http.Error(w, "Only POST supported", http.StatusMethodNotAllowed)
		return
	}

	req.FromValue()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	a.s.WriteDeposit(a.txn)
}
