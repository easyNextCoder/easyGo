package cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"strconv"
	"time"
)

var (
	//New(defaultExpiration, cleanupInterval time.Duration)
	gcache *cache.Cache = cache.New(time.Second*5, time.Millisecond*100)
)

type User struct {
	name string
	age  int
}

func cacheValue() {
	//d = 0  表示使用默认过期时间
	//d = -1 表示永远不过期
	gcache.Set("user1", User{
		name: "kk",
		age:  18,
	}, 0)

	<-time.After(time.Second * 3)
	val, ok := gcache.Get("user1")
	if ok {
		fmt.Printf("loadValue success %+v\n", val.(User))
	} else {
		fmt.Printf("loadValue failed\n")
	}
}

func cachePointer() {

	user := &User{
		name: "kk",
		age:  18,
	}

	gcache.Set("user1", user, 0)

	<-time.After(time.Second * 3)
	val, ok := gcache.Get("user1")
	if ok {
		got, _ := val.(*User)
		got.name = "kkk"
		fmt.Printf("loadPointer success got %+v origin %+v\n", got, user)
	} else {
		fmt.Printf("loadPointer failed\n")
	}
}

func cacheItemCountDelete() {

	count := 0
	for count < 100 {
		count++
		gcache.Set(strconv.Itoa(count), count, 0)
	}

	num := gcache.ItemCount()

	fmt.Printf("cache itemCount %d\n", num)

	count = 0
	for count < 100 {
		count++
		gcache.Delete(strconv.Itoa(count))
	}

	num = gcache.ItemCount()

	fmt.Printf("cache itemCount after Delete %d\n", num)
}
