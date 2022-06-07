package pubsub

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/whosonfirst/go-webhookd/v3/dispatcher"
	"log"
	"testing"
	"time"
)

func TestPubSubDispatcher(t *testing.T) {

	ctx := context.Background()

	redis_host := "localhost:6379"
	redis_channel := "example"

	rdb := redis.NewClient(&redis.Options{
		Addr: redis_host,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		log.Printf("Failed to ping Redis server, %v", err)
		t.Skip()
	}

	hello_world := []byte("hello world")

	dispatcher_uri := fmt.Sprintf("pubsub://%s/%s", redis_host, redis_channel)

	done_ch := make(chan bool)

	go func() {

		defer func() {
			done_ch <- true
		}()

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pubsub := rdb.Subscribe(ctx, redis_channel)

		message_ok := false

		for {

			msg, err := pubsub.ReceiveMessage(ctx)

			if err != nil {
				log.Fatalf("Failed to receive pubsub message, %v", err)
			}

			if msg.Payload == string(hello_world) {
				message_ok = true
				break
			}
		}

		if !message_ok {
			log.Fatalf("Did not receive pubsub message")
		}

		return
	}()

	d, err := dispatcher.NewDispatcher(ctx, dispatcher_uri)

	if err != nil {
		t.Fatalf("Failed to create new dispatcher, %v", err)
	}

	err2 := d.Dispatch(ctx, hello_world)

	if err2 != nil {
		t.Fatalf("Failed to dispatch message, %v", err2)
	}

	<-done_ch
}
