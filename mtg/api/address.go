package api

import (
	"net/http"
)

// Handle Address creation
func AddressHandler(w http.ResponseWriter, req *http.Request) {
	println(req)
}
