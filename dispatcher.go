package pubsub

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/whosonfirst/go-webhookd/v3"
	"github.com/whosonfirst/go-webhookd/v3/dispatcher"
	"net/url"
)

func init() {

	ctx := context.Background()
	err := dispatcher.RegisterDispatcher(ctx, "pubsub", NewPubSubDispatcher)

	if err != nil {
		panic(err)
	}
}

type PubSubDispatcher struct {
	webhookd.WebhookDispatcher
	client  *redis.Client
	channel string
}

func NewPubSubDispatcher(ctx context.Context, uri string) (webhookd.WebhookDispatcher, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	endpoint := u.Host
	channel := u.Path

	client := redis.NewClient(&redis.Options{
		Addr: endpoint,
	})

	// defer client.Close()

	_, err = client.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}

	dispatcher := PubSubDispatcher{
		client:  client,
		channel: channel,
	}

	return &dispatcher, nil
}

func (dispatcher *PubSubDispatcher) Dispatch(ctx context.Context, body []byte) *webhookd.WebhookError {

	select {
	case <-ctx.Done():
		return nil
	default:
		// pass
	}

	rsp := dispatcher.client.Publish(ctx, dispatcher.channel, string(body))

	_, err := rsp.Result()

	if err != nil {

		code := 999
		message := err.Error()

		err := &webhookd.WebhookError{Code: code, Message: message}
		return err
	}

	return nil
}
