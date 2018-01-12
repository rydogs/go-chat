package web

import "net/http"

func Handlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/ws", new(webSocket))
	return mux
}
