package drivers

import "context"

type DriverInterface interface {
	Send(ctx context.Context, queue string, payload string) error
	Listen(queue string, f func(err error, msg string, ack func()))
}
