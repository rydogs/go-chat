package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chat/web"
	"github.com/go-chat/web/middleware"
	"github.com/urfave/negroni"
)

func main() {
	port, newrelicKey := os.Getenv("PORT"), os.Getenv("NEW_RELIC_LICENSE_KEY")

	if port == "" {
		port = "3000"
	}

	n := negroni.Classic() // Includes some default middlewares
	n.Use(middleware.NewRelic("go-chat", newrelicKey))
	n.UseHandler(web.Handlers())

	http.ListenAndServe(":"+port, n)
	fmt.Printf("Server listening on port: %s", port)
}
