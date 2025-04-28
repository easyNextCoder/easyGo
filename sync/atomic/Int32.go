package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func atomicInt32() {
	var m int
	var n atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			n.Add(1)
			m++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("1万个go程每次给原子变量+1，结果:%d\n", n.Load())
	fmt.Printf("1万个go程每次给普通变量+1，结果:%d\n", m)
}

type Info struct {
	name string
	id   int
}

func atomicValue() {

	var v atomic.Value

	var info = Info{
		name: "account info",
		id:   1,
	}

	v.Store(info)

	var wg sync.WaitGroup

	wg.Add(2000)

	for i := 0; i < 1000; i++ {
		go func() {
			vl := v.Load()
			vlt, ok := vl.(Info)
			if ok {
				vlt.id++
				v.Store(vlt)
			}
			wg.Done()
		}()

		go func() {
			vl := v.Load()
			vlt, ok := vl.(Info)
			if ok {
				vlt.id--
				v.Store(vlt)
			}
			wg.Done()
		}()

	}

	vl, _ := v.Load().(Info)

	fmt.Printf("vl is %+v\n", vl)

	wg.Wait()
	time.Sleep(time.Second * 5)
}

func atomicValue2() {
	//atomic.Value只能存储一种类型的变量，否则会panic
	var container atomic.Value

	var info Info
	var infoPointer = new(Info)

	container.Store(info)
	container.Store(infoPointer)
}
