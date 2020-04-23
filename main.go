package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
)

func display(event cloudevents.Event) {
	log.Print(event.String())
}

func main() {
	client, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("could not create client: %v\n", err)
	}

	ctx := context.Background()

	log.Printf("server started on port %d", 8080)

	err = client.StartReceiver(ctx, display)
	if err != nil {
		log.Fatalf("failed to start server, %v", err)
	}
}
