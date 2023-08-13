package api

import (
	"net/http"

	"github.com/dgraph-io/badger/v4"
	"github.com/mvg-fi/common/web"
	"github.com/mvg-fi/mvg-bridge/store"
)

type API struct {
	s   *store.BadgerStore
	txn *badger.Txn
	c   *web.Configuration
}

func NewAPIWorker(s *store.BadgerStore, txn *badger.Txn, c *web.Configuration) *API {
	return &API{
		s:   s,
		txn: txn,
		c:   c,
	}
}

func (a *API) Loop() {
	a.loop(a.c.Host, a.c.Port)
}

func (a *API) loop(host, port string) {
	http.HandleFunc("/api/address", AddressHandler)
	http.HandleFunc("/api/status", StatusHandler)
	http.ListenAndServe(host+":"+port, nil)
}
