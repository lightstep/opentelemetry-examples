package main

import (
	"fmt"
	"github.com/go-stomp/stomp"
)

func Producer(c, quit chan string, conn *stomp.Conn) {
	for {
		select {
		case c <- "msg sent":
			err := conn.Send(
				"/queue/myqueue1",         //destination
				"text/plain",              //content-type
				[]byte("Test message #1")) //body
			if err != nil {
				fmt.Println(err)
				return
			}
		case <-quit:
			fmt.Println("finish")
			return
		}
	}
}

//Connect to ActiveMQ and produce messages
func main() {
	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		fmt.Println(err)
	}

	c := make(chan string)
	quit := make(chan string)
	go Producer(c, quit, conn)

	for {
		fmt.Println(<-c)
	}
	quit <- "read"
}
