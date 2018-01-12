package web

import (
	"net/http"
)

type webSocket struct{}

func (ws *webSocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Sec-WebSocket-Extensions", "application/websocket-events")
	w.Header().Add("Content-Type", "grip")
	w.WriteHeader(200)
}
