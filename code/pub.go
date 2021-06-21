package main

import (
	"context"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"os"
	"time"
)

var (
	// set the environment as instructions.
	pubsubName = os.Getenv("DAPR_PUBSUB_NAME")
	topicName  = "neworder"
)

func main() {
	ctx := context.Background()
	data := []byte("ping")

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	for i := 0; i < 1000; i++ {
		if err := client.PublishEvent(ctx, pubsubName, topicName, data); err != nil {
			panic(err)
		}
		fmt.Println("data published")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Done (CTRL+C to Exit)")
}
