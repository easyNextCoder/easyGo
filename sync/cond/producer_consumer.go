package cond

import (
	"fmt"
	"sync"
	"time"
)

type BlockingQueue struct {
	queue    []int
	capacity int
	lock     sync.Mutex
	notEmpty *sync.Cond
	notFull  *sync.Cond
}

// 构造函数
func NewBlockingQueue(capacity int) *BlockingQueue {
	bq := &BlockingQueue{
		queue:    make([]int, 0),
		capacity: capacity,
	}
	bq.notEmpty = sync.NewCond(&bq.lock)
	bq.notFull = sync.NewCond(&bq.lock)
	return bq
}

// 入队：满了就等待
func (bq *BlockingQueue) Enqueue(item int) {
	bq.lock.Lock()
	defer bq.lock.Unlock()

	// 如果队列满了，等待 notFull 条件
	for len(bq.queue) == bq.capacity {
		fmt.Println("Enqueue wait")
		bq.notFull.Wait()
		fmt.Println("Enqueue wait done")
	}

	// 添加元素
	bq.queue = append(bq.queue, item)

	// 通知等待“非空”的消费者
	bq.notEmpty.Signal()
}

// 出队：空了就等待
func (bq *BlockingQueue) Dequeue() int {
	bq.lock.Lock()
	defer bq.lock.Unlock()

	// 如果队列空，等待 notEmpty 条件
	for len(bq.queue) == 0 {
		fmt.Println("Dequeue wait")
		bq.notEmpty.Wait()
		fmt.Println("Dequeue wait done")
	}

	// 取出队头元素
	item := bq.queue[0]
	bq.queue = bq.queue[1:]

	// 通知等待“非满”的生产者
	bq.notFull.Signal()

	return item
}

func ConsumeProduce() {
	bq := NewBlockingQueue(3)

	// 启动 1 个生产者
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("生产：", i)
			bq.Enqueue(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// 启动 2 个消费者
	for j := 0; j < 2; j++ {
		go func(id int) {
			for {
				item := bq.Dequeue()
				fmt.Printf("消费者 %d 消费了: %d\n", id, item)
				time.Sleep(300 * time.Millisecond)
			}
		}(j)
	}

	time.Sleep(time.Second * 6)
	//<-make(chan int)
	//select {} // 阻塞主协程
}
