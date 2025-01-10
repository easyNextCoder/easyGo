package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func Test_signal_stop(t *testing.T) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT)
	go func() {
		_, ok := <-sigs
		if ok {
			fmt.Println("received sigInt")
			done <- true
		}
	}()

	<-done
}
