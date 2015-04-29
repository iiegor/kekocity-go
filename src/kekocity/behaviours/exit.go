package behaviours

import (
    "os"
    "os/signal"
    "syscall"
)

type hooks []func()

func (h *hooks) Register(f func()) {
	*h = append(*h, f)
}

// Hooks to be run before exiting.
var Hooks = make(hooks, 0)

func Exit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c

		// Run hooks
		for _, hook := range Hooks {
			hook()
		}

		os.Exit(1)
	}()
}
