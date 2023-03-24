package jetstream

import (
	"github.com/demig00d/order-service/config"
	"github.com/nats-io/nats.go"
)

// JetStream -.
type JetStream struct {
	nats.JetStreamContext
	StreamName, StreamSubjects string
	SubjectNameOrderCreated    string
}

func New(cfg config.JetStream) (*JetStream, error) {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	jsCtx, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		return nil, err
	}

	return &JetStream{
		jsCtx,
		cfg.StreamName,
		cfg.StreamSubjects,
		cfg.SubjectNameOrderCreated,
	}, nil
}

func (js *JetStream) CreateStream() error {
	stream, err := js.StreamInfo(js.StreamName)

	// stream not found, create it
	if stream == nil {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     js.StreamName,
			Subjects: []string{js.StreamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (js *JetStream) Close() {
	if js != nil {
		js.Close()
	}
}
