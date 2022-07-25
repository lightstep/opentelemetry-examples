package main

import (
	"fmt"
	"math/rand"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitMQClient struct {
	channelNum int64

	conn     *amqp.Connection
	channels []*amqp.Channel
}

func newRabbitMCQClient(channelNum int) *rabbitMQClient {
	rabbitMQClient := new(rabbitMQClient)
	rabbitMQClient.channels = make([]*amqp.Channel, 30)
	rabbitMQClient.channelNum = rand.Int63n(30)
	var err error
	rabbitMQClient.conn, err = amqp.DialConfig("amqp://user:userpass@localhost:5672/", amqp.Config{Heartbeat: 0})
	if err != nil {
		panic(err)
	}
	return rabbitMQClient
}

func (rmc *rabbitMQClient) channelLength() int {
	return len(rmc.channels)
}

func (rmc *rabbitMQClient) closeConnection() {
	rmc.conn.Close()
}

func (rmc *rabbitMQClient) sendMsg(queue amqp.Queue) {
	err := rmc.channels[0].Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("test"),
		})
	if err != nil {
		panic(err)
	}
}

func main() {
	var clients []*rabbitMQClient
	for i := 0; i < 1; i++ {
		clients = append(clients, newRabbitMCQClient(2))
	}

	client := newRabbitMCQClient(2)
	ch, err := client.conn.Channel()
	if err != nil {
		panic(err)
	}

	queue, _ := ch.QueueDeclare(
		"test", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wai2t
		nil,    // arguments
	)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	body := []byte("test")
	wg.Add(1)
	go func() {
		for {
			fmt.Println("Sending message to rabbiMQ producer:", "test")
			err = ch.Publish(
				"",         // exchange
				queue.Name, // routing key
				false,      // mandatory
				false,      // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        body,
				})
		}
	}()

	client2 := newRabbitMCQClient(2)
	ch2, err := client2.conn.Channel()
	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go func() {
		for {
			fmt.Println("Sending message to rabbiMQ producer:", "test")
			err = ch2.Publish(
				"",         // exchange
				queue.Name, // routing key
				false,      // mandatory
				false,      // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        body,
				})
		}
	}()

	client3 := newRabbitMCQClient(2)
	ch3, err := client3.conn.Channel()
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go func() {
		for {
			fmt.Println("Sending message to rabbiMQ producer:", "test")
			err = ch3.Publish(
				"",         // exchange
				queue.Name, // routing key
				false,      // mandatory
				false,      // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        body,
				})
		}
	}()
	wg.Wait()
}
