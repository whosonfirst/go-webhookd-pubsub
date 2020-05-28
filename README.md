# go-webhookd-pubsub

## Important

Work in progress.

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