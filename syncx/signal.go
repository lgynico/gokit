package syncx

import (
	"os"
	"os/signal"
	"syscall"
)

type Signal struct {
	signalC chan os.Signal
}

func NewSignal() *Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return &Signal{
		signalC: sig,
	}
}

func (p *Signal) Wait() os.Signal {
	return <-p.signalC
}

func (p *Signal) Channel() <-chan os.Signal {
	return p.signalC
}
