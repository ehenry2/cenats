package pub

import (
	"context"
	cenats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/protocol"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"time"
)

const jsonContentType = "application/json"


func NewNatsSender(subject, natsURL string) (*protocol.Sender, error) {
	var s protocol.Sender
	s, err := cenats.NewSender(natsURL, subject, cenats.NatsOptions())
	return &s, err
}

type Handler struct {
	client cloudevents.Client
	payloadReader io.Reader
}

func (h *Handler) Handle(ctx context.Context, eventType, eventSource string) error {
	// read in the payload data
	b, err := ioutil.ReadAll(h.payloadReader)
	if err != nil {
		return err
	}
	// create the event
	e := cloudevents.NewEvent()

	// set metadata
	e.SetID(uuid.New().String())
	e.SetType(eventType)
	e.SetTime(time.Now())
	e.SetSource(eventSource)

	// set the payload
	err = e.SetData(jsonContentType, b)
	if err != nil {
		return err
	}

	// send the event
	return h.client.Send(ctx, e)
}

func NewHandler(sender *protocol.Sender, reader io.Reader) (*Handler, error) {
	client, err := cloudevents.NewClient(*sender)
	if err != nil {
		return nil, err
	}
	return &Handler{client, reader}, nil
}
