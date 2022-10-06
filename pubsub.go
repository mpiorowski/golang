package utils

import (
	"context"
	"log"
	"cloud.google.com/go/pubsub"
	"github.com/gin-gonic/gin"
)

func PublishPubSub(projectId, topicId, msg string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Printf("pubsub.NewClient: %v", err)
		return err
	}
	defer client.Close()

	t := client.Topic(topicId)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		log.Printf("result.Get: %v", err)
		return err
	}
	log.Printf("Published a message with msg id: %v", id)
	return nil
}

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}
func SubscribePubSub(c* gin.Context) (PubSubMessage, error) {
	var message PubSubMessage
	err := c.BindJSON(&message)
	if err != nil {
		log.Printf("c.BindJSON: %v", err)
		return message, err
	}
  log.Printf("Got message: %v", string(message.Message.Data))
  return message, nil
}
