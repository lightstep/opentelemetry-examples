package main

// NOTE: this is a demo application to generate traffic for istio observation. this program as is does not correctly handle memory leaks, does not use proper declartive and object oriented programming golang practices, or tcp connection reconnects

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

var defaultPeerRedialFrequency time.Duration = 2

var peerRedialFrequency = flag.Int("peerRedialFrequency", 30, "the time in second to retry a peer")
var peerURLOpt = flag.String("peerURL", "", "write peer url")
var listenURLOpt = flag.String("listenURL", "", "url to listen on")

type message struct {
	Message string `yaml:"message"`
	Number  int    `yaml:"number"`
}

func main() {
	err := machine()
	if err != nil {
		fmt.Println(err.Error(), "\nexiting the machine")
	}
}

func machine() error {
	hn, _ := os.Hostname()
	fmt.Println("Starting machine:", os.Getenv("NAME"), hn)
	flag.Parse()
	peerRedialFrequencyDuration := defaultPeerRedialFrequency * time.Second
	if *peerRedialFrequency != 0 {
		peerRedialFrequencyDuration = time.Duration(*peerRedialFrequency) * time.Second
	}
	peerURL, ok := os.LookupEnv("PEER_URL")
	listenURL, ok := os.LookupEnv("LISTEN_URL")
	if *peerURLOpt != "" {
		peerURL = *peerURLOpt
	}
	if *listenURLOpt != "" {
		listenURL = *listenURLOpt
	}
	if !ok {
		fmt.Println("ERROR: you must set environmental variables: PEER_URL && LISTEN_URL")
		os.Exit(1)
	}
	errs := make(chan error, 1)
	bytes := make(chan []byte, 2)
	ln, err := net.Listen("tcp", listenURL)
	if err != nil {
		fmt.Println("Failed to start machine:", err)
		errs <- err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func(chan<- []byte) {
		var once sync.Once
		for {
			conn, err := ln.Accept()
			if err != nil {
				errs <- err
			}
			defer func() {
				fmt.Println("Closing listener:", ln.Addr(), "on", os.Getenv("Name"))
				ln.Close()
				wg.Done()
				close(bytes)
			}()
			if conn != nil {
				go func() {
					seed := func() {
						if os.Getenv("SEED") == "true" {
							fmt.Println("Starting message chain with seed")
							time.Sleep(1 * time.Minute)
							msg := new(message)
							msg.Message = "Start message"
							msg.Number = 0
							b, err := yaml.Marshal(msg)
							if err != nil {
								fmt.Println(err)
							}
							bytes <- b
						}
					}
					once.Do(seed)
					defer func() {
						fmt.Println("Closing connection:", conn.RemoteAddr(), "on listener", os.Getenv("Name"))
						conn.Close()
					}()
					fmt.Println("Accepting messages on", ln.Addr(), "from", conn.RemoteAddr())
					b := make([]byte, 32*8)
					for {
						_, err := conn.Read(b)
						if err != nil {
							continue
						}
						bytes <- b
					}
				}()
			}
		}
	}(bytes)
	wg.Add(1)
	go func(<-chan []byte) {
		peerRecount := 0
		connected := false
		var conn net.Conn
		var writeFrequency = 3 * time.Second
		for {
			if peerRecount <= 5 && connected == false {
				conn, err = net.Dial("tcp", peerURL)
				if err != nil {
					fmt.Println("Unable to dial peer redialing, ", peerRecount, err.Error())
					time.Sleep(peerRedialFrequencyDuration)
					peerRecount++
					continue
				}
				connected = true
				defer func() {
					fmt.Println("Closing connection to", conn.RemoteAddr(), "on", conn.LocalAddr())
					conn.Close()
					wg.Done()
				}()
				fmt.Println("Successfully dialed: ", conn.RemoteAddr())
			}
			if peerRecount >= 5 || conn == nil {
				errs <- errors.New("Failed to connect to peer")
				return
			}
			select {
			case b := <-bytes:
				msg := new(message)
				time.Sleep(writeFrequency)
				yaml.Unmarshal(b, &msg)
				fmt.Println(msg.Message, msg.Number)
				msg.Message = "Passage Message"
				msg.Number++
				b, err := yaml.Marshal(msg)
				if err != nil {
					fmt.Println(err)
					continue
				}
				conn.Write(b)
			default:
				msg, _ := json.Marshal(message{Message: "reset", Number: 0})
				conn.Write(msg)
			}
		}
	}(bytes)
	for {
		select {
		case err := <-errs:
			return err
		}
	}
}
