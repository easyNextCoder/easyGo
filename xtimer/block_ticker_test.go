package xtimer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func Test_block_ticker(t *testing.T) {
	wg.Add(1)
	go func() {
		tick := time.NewTicker(time.Millisecond * 200)
		cnt := 0
		for {
			select {
			case <-tick.C:
				fmt.Printf("do time %v\n", time.Now())

				if cnt == 2 {
					fmt.Printf("delay 4s start time %v \n", time.Now())
					<-time.After(time.Second * 4) //when blocking no one enter, after blocking only one enter at the same time.
					fmt.Printf("delay 4s end time %v\n", time.Now())
				}

				cnt++
				if cnt > 4 {
					wg.Done()
					return
				}
			}
		}

	}()
	wg.Wait()

}
