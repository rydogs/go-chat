package main

import (
	"net/http"
	"fmt"
	"os"

	"github.com/urfave/negroni"
	"github.com/newrelic/go-agent"
)

func main() {
	port, newrelicKey := os.Getenv("PORT"), os.Getenv("NEW_RELIC_LICENSE_KEY")

	if port == "" {
		port = "3000"
	}

	config := newrelic.NewConfig("go-chat", newrelicKey)
	app, err := newrelic.NewApplication(config)
	if (err != nil) {
		fmt.Println("Failed to load new relic agent")
	}

	mux := http.NewServeMux()
	mux.HandleFunc(newrelic.WrapHandleFunc(app, "/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	}))

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	http.ListenAndServe(":" + port, n)
}
