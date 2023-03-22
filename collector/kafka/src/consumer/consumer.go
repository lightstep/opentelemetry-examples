package main

import (
	"context"
	"fmt"
)

const (
	topic     = "test-topic"
	kafkaAddr = "kafka:9092"
	partition = 2
	groupID   = "my-group"
)

// StartConsumer starts consuming messages from the Kafka topic.
func StartConsumer(ctx context.Context) {
	consumer := newKafkaConsumer()
	defer consumer.Close()

	for {
		select {
		case <-ctx.Done():
			return // Exit the function when the context is canceled.
		default:
			message, err := readMessage(ctx, consumer)
			if err != nil {
				// handle error
				fmt.Printf("Handle Error: %s\n", err)
			} else {
				// handle success
				fmt.Printf("Consumed message: %s\n", string(message.Value))
			}
		}
	}
}

// newKafkaConsumer creates a new Kafka consumer instance.
func newKafkaConsumer() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafkaAddr},
		Topic:     topic,
		Partition: partition,
		GroupID:   groupID,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
}

// readMessage reads a message from the Kafka topic.
func readMessage(ctx context.Context, consumer *kafka.Reader) (kafka.Message, error) {
	message, err := consumer.ReadMessage(ctx)
	if err != nil {
		return kafka.Message{}, err
	}
	return message, nil
}
