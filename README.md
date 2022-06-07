# go-webhookd-pubsub

Go package to implement the `whosonfirst/go-webhookd` interfaces for dispatching webhooks messages to a Redis PubSub channel.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/whosonfirst/go-webhookd-pubsub.svg)](https://pkg.go.dev/github.com/whosonfirst/go-webhookd-pubsub)

Before you begin please [read the go-webhookd documentation](https://github.com/whosonfirst/go-webhookd/blob/master/README.md) for an overview of concepts and principles.

## Usage

```
import (
	_ "github.com/go-webhookd-pubsub"
)
```

## Dispatchers

### PubSub

The `PubSub` dispatcher will send messages to a Redis PubSub channel. It is defined as a URI string in the form of:

```
pubsub://{REDIS_HOST}:{REDIS_PORT}/{REDIS_CHANNEL}
```

#### Properties

| Name | Value | Description | Required |
| --- | --- | --- | --- |
| redis_host | string | The Redis host to publish to | yes |
| redis_port | int | The Redis port to publish to | yes |
| redis_channle | string | The name of the Redis channel to publish to | yes |


## See also

* https://github.com/whosonfirst/go-webhookd