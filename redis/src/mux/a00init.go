package mux

import (
	"net/http"
)

func init() {
	go http.ListenAndServe(":32123", GetMuxCli())
}
