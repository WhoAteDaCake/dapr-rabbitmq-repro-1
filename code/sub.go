package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"os"
	"time"
	"math/rand"
)

// Subscription to tell the dapr what topic to subscribe.
// - PubsubName: is the name of the component configured in the metadata of pubsub.yaml.
// - Topic: is the name of the topic to subscribe.
// - Route: tell dapr where to request the API to publish the message to the subscriber when get a message from topic.
var sub = &common.Subscription{
	PubsubName: os.Getenv("DAPR_PUBSUB_NAME"),
	Topic:      "neworder",
	Route:      "/orders",
}

func main() {
	log.Printf("Starting")
	s := daprd.NewService(":" + os.Getenv("APP_PORT"))

	if err := s.AddTopicEventHandler(sub, eventHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	n := rand.Intn(10)
	log.Printf("Sleeping before finishing %d", n)
	time.Sleep(time.Duration(n) * time.Second)
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	return false, nil
}
