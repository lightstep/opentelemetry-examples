package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pong)
	}

	rand.Seed(time.Now().UnixNano())

	keys := []string{"key1", "key2", "key3", "key4", "key5"}
	for {
		key := keys[rand.Intn(len(keys))]
		value := rand.Intn(1000000)
		err := client.Set(key, value, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
