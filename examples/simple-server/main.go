package main

import (
	"net/http"

	"github.com/functionalfoundry/graphqlws"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.Info("Starting example server on :8085")

	// Create subscription manager and GraphQL WS handler
	subscriptionManager := graphqlws.NewSubscriptionManager()
	websocketHandler := graphqlws.NewHandler(graphqlws.HandlerConfig{
		SubscriptionManager: subscriptionManager,
		Authenticate: func(token string) (interface{}, error) {
			return "Default user", nil
		},
	})

	// Serve the GraphQL WS endpoint
	http.Handle("/subscriptions", websocketHandler)
	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.WithField("err", err).Error("Failed to start server")
	}
}
