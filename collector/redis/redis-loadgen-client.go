package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pong)
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Starting load generator...")
	fmt.Println("Setting and getting random keys with 95% hit ratio after warmup.")
	fmt.Println("Press Ctrl-C to stop.")
	for {
		keyToSet := fmt.Sprintf("key%d", rand.Intn(100))
		keyToGet := fmt.Sprintf("key%d", 5+rand.Intn(100))
		value := rand.Intn(1000000)
		err := client.Set(keyToSet, value, 0).Err()
		if err != nil {
			fmt.Println("Error setting key: ", err)
		}
		val, err := client.Get(keyToGet).Result()
		if err != nil {
			fmt.Println("Error getting key: ", err)
		}
		fmt.Println("got key: ", keyToGet, " value: ", val)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
