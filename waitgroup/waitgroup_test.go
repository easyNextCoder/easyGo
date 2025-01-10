package waitgroup

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

//waitgroup 是可以通过函数进行传递的具体使用见signal章节

func Test_waitGroup(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		go func(i int) {
			wg.Add(1)
			for j := 0; j < 10; j++ {
				fmt.Printf("I am go %d work on %d \n", i, j)
				time.Sleep(time.Duration(rand.Intn(1000)))
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
