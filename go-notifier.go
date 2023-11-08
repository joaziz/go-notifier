package go_notifier

import "github.com/joaziz/go-notifier/internal"

func NewInterNotifier(name string) *internal.Notifier {
	return internal.New(name)
}

func NewExternalNotifier() {

}
