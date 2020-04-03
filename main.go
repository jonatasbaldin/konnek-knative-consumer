package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/knative/eventing-sources/pkg/kncloudevents"
)

func display(event cloudevents.Event) {
	log.Print(event.String())
}

func main() {
	client, err := kncloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("could not create client: %v\n", err)
	}

	ctx := context.Background()

	err = client.StartReceiver(ctx, display)
	if err != nil {
		log.Fatalf("failed to start receiver, %v", err)
	}
}
