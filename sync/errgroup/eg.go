package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"time"
)

func ReturnFirstError() {
	//会全部执行，并返回遇到的第一个错误
	var g errgroup.Group
	var urls = []string{
		"https://www.golang.org/",
		"http://www.google.comm/", //第一个错误会被返回
		"http://www.gooooogle.comm/",
		"https://www.baidu.com/",
		"https://www.google.com",
	}
	for i := range urls {
		url := urls[i]
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				bytes, _ := io.ReadAll(resp.Body)
				fmt.Println("get ", url, len(bytes))
				resp.Body.Close()
			}
			return err
		})
	}

	if err := g.Wait(); err == nil {
		fmt.Printf("Successfully fetched all URLs.\n")
	} else {
		fmt.Printf("Failed fetched all URLs. err %v \n", err)
	}
}

func CancelWhenError() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	for i := 0; i < 3; i++ {
		i := i
		g.Go(func() error {
			if i == 1 {
				time.Sleep(time.Second * 1)
				return fmt.Errorf("任务 %d 失败", i)
			}

			select {
			case <-ctx.Done():
				fmt.Printf("任务 %d 被取消了\n", i)
				return ctx.Err()
			}

		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("有任务失败:", err)
	}
}
