package external

import (
	"context"
	"github.com/joaziz/go-notifier/external/drivers"
)

func NewNotifier(drive drivers.DriverInterface) *Notifier {
	return &Notifier{
		drive: drive,
	}
}

type Notifier struct {
	drive drivers.DriverInterface
}

func (n *Notifier) Listen(queue string, f func(err error, msg string, ack func())) {
	n.drive.Listen(queue, f)
}

func (n *Notifier) Send(ctx context.Context, queue string, msg string) error {
	return n.drive.Send(ctx, queue, msg)
}

// -----------------------
