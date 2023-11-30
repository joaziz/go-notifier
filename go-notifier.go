package go_notifier

import (
	"github.com/joaziz/go-notifier/external"
	"github.com/joaziz/go-notifier/external/drivers"
	"github.com/joaziz/go-notifier/local"
	"log/slog"
)

func NewInterNotifier(name string, opt local.Options) *local.Notifier {
	if opt.Logger == nil {
		opt.Logger = slog.Default()
	}
	return local.New(name, opt)
}

func NewExternalNotifier(driver drivers.DriverInterface) *external.Notifier {
	return external.NewNotifier(driver)
}
