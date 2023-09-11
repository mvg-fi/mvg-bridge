package api

import (
	"log"
	"net/http"
	"time"

	"github.com/mvg-fi/common/web"
	"github.com/mvg-fi/go-limiter/httplimit"
	"github.com/mvg-fi/go-limiter/memorystore"
	"github.com/mvg-fi/mvg-bridge/store"
)

type API struct {
	s *store.BadgerStore
	c *web.Configuration
}

func NewAPIWorker(s *store.BadgerStore, c *web.Configuration) *API {
	return &API{
		s: s,
		c: c,
	}
}

func InitIPRateLimit() *httplimit.Middleware {
	store, err := memorystore.New(&memorystore.Config{
		// Number of tokens allowed per interval.
		Tokens: 60,

		// Interval until tokens reset.
		Interval: time.Minute,
	})
	if err != nil {
		log.Fatal(err)
	}
	middleware, err := httplimit.NewMiddleware(store, httplimit.IPKeyFunc())
	if err != nil {
		log.Fatal(err)
	}
	return middleware
}

func (a *API) run(host, port string) {
	//ipRateLimit := InitIPRateLimit()

	http.Handle("/price/simple", a.PriceSimpleHandler())
	http.Handle("/price/all", a.PriceAllHandler())

	//http.Handle("/address", ipRateLimit.Handle(http.Handler(a.AddressHandler(a.s))))
	http.HandleFunc("/status", a.StatusHandler())
	http.ListenAndServe(host+":"+port, nil)
}

func (a *API) Run() {
	a.run(a.c.Host, a.c.Port)
}
