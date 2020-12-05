package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()

	const (
		port1 = ":8889"
		port2 = ":8888"
	)

	// 开启服务
	go StartServer(ctx, port2, "")
	time.Sleep(1 * time.Second)

	go StartServer(ctx, port1, port2)
	time.Sleep(1 * time.Second)

	if err := do(ctx, 10*time.Second, "http://127.0.0.1"+port1); err != nil {
		fmt.Printf("err: %+v\n", err)
	}

	for {
	}
}

func do(ctx context.Context, timeout time.Duration, url string) (err error) {
	// 请求
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Printf("data: %s\n", data)

	return
}
