package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

var registed = make([]chan *sync.WaitGroup, 0, 8)
var registerChan = make(chan chan *sync.WaitGroup)

func Register(p chan *sync.WaitGroup) {
	registerChan <- p
}

func watch() {
	stopSignalChan := make(chan os.Signal, 1)
	signal.Notify(stopSignalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	for {
		select {
		case c, ok := <-registerChan:
			if ok {
				registed = append(registed, c)
			}
		case _, ok := <-stopSignalChan:
			if ok {
				var wg = sync.WaitGroup{}
				wg.Add(len(registed))
				for _, c := range registed {
					c <- &wg
				}
				wg.Wait()
				os.Exit(0)
			}
		}
	}
}

func flushData1() {
	stopSignalChan := make(chan *sync.WaitGroup, 1)
	Register(stopSignalChan)

	for {
		select {
		case wg, ok := <-stopSignalChan:
			if ok {
				fmt.Println("flushData1 start sleep start")
				time.Sleep(time.Second * 5)
				fmt.Println("flushData1 start sleep end")
				wg.Done()
			}
		}
	}
}

func flushData2() {
	stopSignalChan := make(chan *sync.WaitGroup, 1)
	Register(stopSignalChan)

	for {
		select {
		case wg, ok := <-stopSignalChan:
			if ok {
				fmt.Println("flushData2 start sleep start")
				time.Sleep(time.Second * 8)
				fmt.Println("flushData2 start sleep end")
				wg.Done()
			}
		}
	}
}

func Test_signalSaveBeforeStop(t *testing.T) {
	go watch()
	go flushData1()
	go flushData2()
	time.Sleep(time.Second * 20)
	fmt.Println("main done")
}
