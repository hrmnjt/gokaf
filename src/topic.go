package src

import (
	"context"
)

type Topic struct {
	ctx       context.Context
	name      string
	channel   chan messageInterface
	consumers []*consumer
	producer  *producer
}

func NewTopic(ctx context.Context, name string) *Topic {
	channelTopic := make(chan messageInterface)
	return &Topic{
		ctx,
		name,
		channelTopic,
		[]*consumer{},
		newProducer(ctx, &channelTopic),
	}
}

func (t *Topic) AddConsumer() {
	t.consumers = append(t.consumers, newConsumer(t.ctx, &t.channel))
}

func (t *Topic) AddConsumers(num int) {
	for i := 0; i < num; i += 1 {
		t.AddConsumer()
	}
}

func (t *Topic) Publish(message messageInterface) {
	t.producer.publish(message)
}

func (t *Topic) Run() {
	if len(t.consumers) == 0 { t.AddConsumer() }
	for _, c := range t.consumers {
		c.run()
	}
}
