package main

import (
	"fmt"
	"github.com/go-stomp/stomp"
)

//Connect to ActiveMQ and listen for messages
func main() {
	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		fmt.Println(err)
	}

	sub, err := conn.Subscribe("/queue/myqueue1", stomp.AckAuto)
	if err != nil {
		fmt.Println(err)
	}
	for {
		msg := <-sub.C
		fmt.Println(msg)
	}

	err = sub.Unsubscribe()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Disconnect()
}
