package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"

	// Imports the Google Cloud Pub/Sub client package
	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()

	// Your Google Cloud Platform project ID
	projectID := "project-id"

	// Creates a client
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// The name for the new topic
	topicName := "test-topic"

	// Creates the new topic
	topic, err := client.CreateTopic(ctx, topicName)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	fmt.Printf("Topic %v created.", topic)
}