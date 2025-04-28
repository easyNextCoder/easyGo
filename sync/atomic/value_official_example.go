package atomic

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"
)

var xv int = 0
var controls []chan int = []chan int{}

func loadConfig() map[string]string {
	xv++
	return map[string]string{strconv.Itoa(xv): strconv.Itoa(xv + 1)}
}

func requests() chan int {
	res := make(chan int)
	controls = append(controls, res)
	return res
}

func valueOfficialExampleWork() {
	var config atomic.Value // holds current server configuration
	// Create initial config value and store into config.
	config.Store(loadConfig())
	go func() {
		// Reload config every 10 seconds
		// and update config value with the new version.
		for {
			time.Sleep(10 * time.Millisecond)
			config.Store(loadConfig())
		}
	}()
	// Create worker goroutines that handle incoming requests
	// using the latest config value.
	for i := 0; i < 10; i++ {
		go func(idx int) {
			for r := range requests() {
				//load latest config
				c := config.Load()
				// Handle request r using config c.
				fmt.Printf("do request %v use latest config %v\n", r, c)
			}
			fmt.Printf("goroutine idx %d exit\n", idx)
		}(i)
	}

	for i := 0; i < len(controls); i++ {
		fmt.Println("sending signal", i)
		<-time.After(time.Millisecond * time.Duration(rand.Intn(100)))
		controls[i] <- i
		close(controls[i])
	}

	time.Sleep(time.Second * 3)

}
