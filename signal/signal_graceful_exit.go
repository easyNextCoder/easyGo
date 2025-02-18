package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var register = make([]chan *sync.WaitGroup, 0, 8)
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
				register = append(register, c)
			}
		case _, ok := <-stopSignalChan:
			if ok {
				var wg = sync.WaitGroup{}
				wg.Add(len(register))
				for _, c := range register {
					c <- &wg
				}
				wg.Wait()
				os.Exit(0) //正式运行的系统需要加
			}
		}
	}
}

func serverModule1() {
	stopSignalChan := make(chan *sync.WaitGroup, 1)
	Register(stopSignalChan)

	for {
		select {
		case wg, ok := <-stopSignalChan:
			if ok {
				time.Sleep(time.Second * 1)
				fmt.Println("server1 saving cache before system shutdown")
				time.Sleep(time.Second * 4)
				wg.Done()
			}
		}
	}
}

func serverModule2() {
	stopSignalChan := make(chan *sync.WaitGroup, 1)
	Register(stopSignalChan)

	for {
		select {
		case wg, ok := <-stopSignalChan:
			if ok {
				time.Sleep(time.Second * 2)
				fmt.Println("server2 saving cache before system shutdown")
				time.Sleep(time.Second * 5)
				wg.Done()
			}
		}
	}
}

func main() {
	go watch()
	go serverModule1()
	go serverModule2()
	time.Sleep(time.Second * 60)
	fmt.Println("system shutdown")
}
