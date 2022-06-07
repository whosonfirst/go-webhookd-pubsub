package pubsub

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/whosonfirst/go-webhookd/v3"
	"github.com/whosonfirst/go-webhookd/v3/dispatcher"
	_ "log"
	"net/url"
	"strings"
)

func init() {

	ctx := context.Background()
	err := dispatcher.RegisterDispatcher(ctx, "pubsub", NewPubSubDispatcher)

	if err != nil {
		panic(err)
	}
}

// PubSubDispatcher implements the `webhookd.WebhookDispatcher` interface for dispatching messages to a Redis PubSub channel.
type PubSubDispatcher struct {
	webhookd.WebhookDispatcher
	// client is a `redis.Client` instance used to deliver messages.
	client *redis.Client
	// channel is the name of the Redis PubSub channel to publish messages to.
	channel string
}

// NewPubSubDispatcher returns a new `PubSubDispatcher` instance configured by 'uri' in the form of:
//
// 	pubsub://{REDIS_HOST}:{REDIS_PORT}/{REDIS_CHANNEL}
func NewPubSubDispatcher(ctx context.Context, uri string) (webhookd.WebhookDispatcher, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	endpoint := u.Host
	channel := u.Path

	channel = strings.TrimLeft(channel, "/")

	client := redis.NewClient(&redis.Options{
		Addr: endpoint,
	})

	// defer client.Close()

	_, err = client.Ping(ctx).Result()

	if err != nil {
		return nil, fmt.Errorf("Failed to ping Redis host, %w", err)
	}

	d := PubSubDispatcher{
		client:  client,
		channel: channel,
	}

	return &d, nil
}

// Dispatch() relays 'body' to the Redis PubSub channel defined when 'd' was instantiated.
func (d *PubSubDispatcher) Dispatch(ctx context.Context, body []byte) *webhookd.WebhookError {

	select {
	case <-ctx.Done():
		return nil
	default:
		// pass
	}

	rsp := d.client.Publish(ctx, d.channel, string(body))

	_, err := rsp.Result()

	if err != nil {

		code := 999
		message := err.Error()

		err := &webhookd.WebhookError{Code: code, Message: message}
		return err
	}

	return nil
}
