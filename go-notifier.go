package go_notifier

import (
	"github.com/joaziz/go-notifier/external"
	"github.com/joaziz/go-notifier/external/drivers"
	"github.com/joaziz/go-notifier/internal"
	"log/slog"
)

func NewInterNotifier(name string, opt internal.Options) *internal.Notifier {
	if opt.Logger == nil {
		opt.Logger = slog.Default()
	}
	return internal.New(name, opt)
}

func NewExternalNotifier(driver drivers.DriverInterface) *external.Notifier {
	return external.NewNotifier(driver)
}
