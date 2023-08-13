package api

import (
	"net/http"

	"github.com/mvg-fi/common/web"
)

type API struct {
	c *web.Configuration
}

func NewAPIWorker(c *web.Configuration) *API {
	return &API{
		c: c,
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
