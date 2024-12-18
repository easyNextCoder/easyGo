package cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	gcache *cache.Cache = cache.New(time.Second*5, time.Millisecond*100)
)

type User struct {
	name string
	age  int
}

func work() {
	//d = 0表示使用默认过期时间
	gcache.Set("user1", User{
		name: "kk",
		age:  18,
	}, 0)

	//d = -1表示永远不过期

	loadData()
}

func loadData() {

	<-time.After(time.Second * 3)
	val, ok := gcache.Get("user1")
	if ok {
		fmt.Printf("loadData success %+v\n", val.(User))
	} else {
		fmt.Printf("loadData failed\n")
	}
}
