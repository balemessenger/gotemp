package pkg

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type SignalDef struct {
}

var Signal = SignalDef{}

func (SignalDef) Wait() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signals
		fmt.Println("Signal received: ", sig)
		done <- true
	}()
	<-done
}
