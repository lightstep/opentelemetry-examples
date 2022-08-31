package main

import (
	"fmt"
    "math/rand"
	"net"
)

func respond(conn *net.UDPConn, addr *net.UDPAddr) {
    var err error
    // 2/3 times responds "up" and rest "down"
    if (rand.Intn(3) != 0) {
	    _, err = conn.WriteToUDP([]byte("up"), addr)
    } else {
	    _, err = conn.WriteToUDP([]byte("down"), addr)
    }
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 9876,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return

	}
	for {
		_, remoteaddr, err := conn.ReadFromUDP(p)
		if err != nil {
            fmt.Printf("Could not read. remoteaddr: %v - %v", remoteaddr, err)
			continue
		}
		go respond(conn, remoteaddr)
	}
}
