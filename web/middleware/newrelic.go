package middleware

import (
	"fmt"
	"net/http"

	"github.com/newrelic/go-agent"
)

type newRelic struct {
	App newrelic.Application
}

func NewRelic(appName string, licenseKey string) *newRelic {
	config := newrelic.NewConfig(appName, licenseKey)
	if licenseKey == "" {
		config.Enabled = false
	}
	app, err := newrelic.NewApplication(config)
	if err != nil {
		fmt.Printf("Failed to load new relic agent: %s", err)
	}
	return &newRelic{app}
}

func (n *newRelic) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	txn := n.App.StartTransaction(r.RequestURI, rw, r)
	defer txn.End()
	next(rw, r)
}
