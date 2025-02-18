[代码地址](https://github.com/easyNextCoder/easyGo/blob/main/signal/signal_graceful_exit.go)

## 用法: 实现优雅退出

* signal信号产生
  * SIGINT信号（Ctrl + C），表示中断，默认行为就是终止程序
  * SIGQUIT信号（Ctrl + \），跟sigint信号差不多，但这个信号会生成core文件，同时在终端打印日志
  * SIGTERM信号（kill pid），通常supervisorctl stop xxx 会发出此信号
  * SIGKILL信号（kill -9 pid），程序无法捕获会立即执行退出，无论此信号是否被监听都会立即退出程序

通常只监听SIGINT、SIGQUIT、SIGTERM这三个信号，然后通知各个在此注册的模块保存缓存信息


```go
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
				//os.Exit(0) //正式运行的系统需要加
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

func Test_signalSaveBeforeStop(t *testing.T) {
	go watch()
	go serverModule1()
	go serverModule2()
	time.Sleep(time.Second * 12)
	fmt.Println("system shutdown")
}

```

