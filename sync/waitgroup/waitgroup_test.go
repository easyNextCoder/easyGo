package waitgroup

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

//waitGroup 是可以通过函数进行传递的具体使用见signal章节

func Test_waitGroup(t *testing.T) {

	var wg sync.WaitGroup

	// Typically this means the calls to Add should execute before the statement
	// creating the goroutine or other event to be waited for.
	// 最好在开go程之前设置完毕，否则可能出现add没有设置完毕，wait已经执行了
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 5; j++ {
				fmt.Printf("I am go %d work on %d \n", i, j)
				time.Sleep(time.Duration(rand.Intn(1000)))
			}
			//自己执行完毕通知
			wg.Done()
		}(i)
	}

	//等10个go程都执行完毕再退出
	wg.Wait()
}
