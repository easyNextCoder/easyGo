package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"strconv"
	"testing"
	"time"
)

var client *clientv3.Client

func init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"0.0.0.0:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("etcd client init err %s\n", err)
		log.Fatal(err)
	}

	client = cli
}

func (r *Room) get(key string) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	resp, err := client.Get(ctx, "mykey")
	cancel()
	if err != nil {
		fmt.Printf("room get err %s, room leaseId %d\n", err, r.leaseId)
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("key:%s, value:%+v\n", ev.Key, ev.Value)
	}
	return nil
}

func (r *Room) put(key string, value string) {
	_, err := client.Put(context.Background(), key, value, clientv3.WithLease(r.leaseId))
	if err != nil {
		fmt.Printf("room leaseId %d, put err %s\n", r.leaseId, err)
	}
}

type Room struct {
	leaseId clientv3.LeaseID
}

func newRoom() *Room {

	leaseId := newLease()
	r := &Room{leaseId: leaseId}

	go func() {
		maxNum := 100
		cnt := 0
		tick := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-tick.C:
				cnt++
				if cnt%2 == 0 {
					r.put("/z/g/l/r_"+strconv.FormatInt(int64(leaseId), 10)+strconv.Itoa(cnt%maxNum), strconv.Itoa(cnt))
				} else {
					r.get("/z/g/l/r_" + strconv.FormatInt(int64(leaseId), 10) + strconv.Itoa(cnt%maxNum))
				}
			}
		}
	}()

	return r
}

func newLease() clientv3.LeaseID {

	lease, err := client.Grant(context.Background(), 10)
	if err != nil {
		fmt.Printf("client newLease err %s\n", err)
		log.Fatal(err)
	}

	keepAliveCtx, keepAliveCancel := context.WithCancel(context.Background())

	ch, err := client.KeepAlive(keepAliveCtx, lease.ID)
	if err != nil {
		fmt.Printf("client keepAlive err %s\n", err)
		keepAliveCancel()
		return -1
	}

	go func() {
		for {
			select {
			case aliveRsp := <-ch:
				if aliveRsp == nil {
					fmt.Printf("aliveRsp is nil\n")
					keepAliveCancel()
					return
				}
			}
		}
	}()

	return lease.ID
}

func Test_performance(t *testing.T) {
	i := 0
	for i < 3000 {
		i++
		fmt.Printf("start idx %d room\n", i)
		newRoom()
	}

	select {}
}
