package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "test-topic"
	kafkaAddr     = "kafka:9092"
	minMessageLen = 10
	maxMessageLen = 100
)

// StartProducer starts producing random messages to the Kafka topic.
func StartProducer(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	producer := newKafkaProducer()
	defer producer.Close()

	for {
		select {
		case <-ctx.Done():
			return // Exit the goroutine when the context is canceled.
		default:
			message := randomMessage()
			err := writeMessage(ctx, producer, message)
			if err != nil {
				// handle error
				fmt.Printf("Handle Error: %s\n", err)
			} else {
				// handle success
				fmt.Printf("Produced message: %s\n", string(message))
			}
			time.Sleep(1 * time.Second)
		}
	}
}

// newKafkaProducer creates a new Kafka producer instance.
func newKafkaProducer() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaAddr),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// writeMessage writes a message to the Kafka topic.
func writeMessage(ctx context.Context, producer *kafka.Writer, message []byte) error {
	return producer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   []byte(strconv.Itoa(rand.Int())),
			Value: message,
		},
	)
}

// randomMessage generates a random string message.
func randomMessage() []byte {
	length := rand.Intn(maxMessageLen-minMessageLen) + minMessageLen
	message := make([]byte, length)
	for i := 0; i < length; i++ {
		message[i] = byte(rand.Intn(26) + 97)
	}
	return message
}
