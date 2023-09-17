package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/mvg-fi/go-limiter/httplimit"
	"github.com/mvg-fi/go-limiter/memorystore"
	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/store"
	"github.com/mvg-fi/mvg-bridge/users"
)

type API struct {
	s *store.BadgerStore
	c *config.Configuration
}

func NewAPIWorker(s *store.BadgerStore, c *config.Configuration) *API {
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

func (a *API) run(ctx context.Context, host, port string) {
	proxy := users.NewProxy(ctx, a.c)
	ipRateLimit := InitIPRateLimit()

	http.Handle("/price/simple", a.PriceSimpleHandler())
	http.Handle("/price/all", a.PriceAllHandler())
	http.Handle("/order/new", ipRateLimit.Handle(http.Handler(a.OrderHandler(ctx, proxy, a.s))))
	http.HandleFunc("/status", a.StatusHandler())
	http.ListenAndServe(host+":"+port, nil)
}

func (a *API) Run(ctx context.Context) {
	a.run(ctx, a.c.API.Host, a.c.API.Port)
}
