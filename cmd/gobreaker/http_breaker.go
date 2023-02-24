package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	breaker "github.com/sony/gobreaker"
)

func main() {
	// reference: https://github.com/sony/gobreaker/blob/master/example/http_breaker.go

	// 创建一个新的 CircuitBreaker 实例。
	breaker := breaker.NewCircuitBreaker(breaker.Settings{
		Name:        "example",
		MaxRequests: 2,
		//Interval:    breaker.DefaultInterval,
		//Timeout:     breaker.DefaultTimeout,
		OnStateChange: func(name string, from breaker.State, to breaker.State) {
			fmt.Printf("Circuit breaker '%s' changed state from %s to %s\n", name, from, to)
		},
	})

	// 封装 http 请求，实现熔断逻辑。
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := breaker.Execute(func() (interface{}, error) {
			resp, err := http.Get("https://example.com/")
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			return resp.Body, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		respBody := body.(io.ReadCloser)
		defer respBody.Close()

		io.Copy(w, respBody)
	})

	// 启动 HTTP 服务器并监听端口。
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
