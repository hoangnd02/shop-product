package services

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaClient struct {
	client *kgo.Client
}

func NewKafkaClient() *KafkaClient {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.AllowAutoTopicCreation(),
	)
	if err != nil {
		panic(err)
	}

	return &KafkaClient{
		client: cl,
	}
}

func (c *KafkaClient) Produce(record *kgo.Record) {
	c.client.Produce(context.Background(), record, func(r *kgo.Record, err error) {
		if err != nil {
			panic(err)
		}

		fmt.Println(string(r.Value))
	})
}
