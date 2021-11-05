package sub

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/protocol"
	cenats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	"log"
)

func NewNatsReceiver(natsURL, subject string) (*protocol.Receiver, error) {
	var receiver protocol.Receiver
	receiver, err := cenats.NewConsumer(natsURL, subject, cenats.NatsOptions())
	if err != nil {
		return nil, err
	}
	return &receiver, nil
}

type Handler struct {
	client cloudevents.Client
}

func (h *Handler) Handle() {
	log.Fatal(
		h.client.StartReceiver(context.Background(), func (ctx context.Context, event cloudevents.Event) {
			log.Println("Event received: ")
			log.Println(string(event.Data()))
		}))
}

func NewHandler(receiver *protocol.Receiver) (*Handler, error) {
	client, err := cloudevents.NewClient(*receiver)
	if err != nil {
		return nil, err
	}
	return &Handler{client}, nil
}
