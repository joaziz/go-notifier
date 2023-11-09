package drivers

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type RedisDriver struct {
	client *redis.Client
}

func NewRedisDriver(opt *redis.Options) *RedisDriver {
	var redisClient = redis.NewClient(opt)
	return &RedisDriver{
		client: redisClient,
	}
}

func (r RedisDriver) Send(ctx context.Context, queue string, payload string) error {

	if err := r.client.Publish(ctx, queue, payload).Err(); err != nil {
		return err
	}
	return nil
}

func (r RedisDriver) Listen(queue string, f func(err error, msg string, ack func())) {
	ctx := context.Background()
	subscriber := r.client.Subscribe(ctx, queue)

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			f(err, "", func() {})
			return
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")

		f(nil, msg.Payload, func() { fmt.Println("Ack") })

	}
}
