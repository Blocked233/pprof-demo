package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	for {
		go leakGoroutine()
		time.Sleep(10 * time.Millisecond)
	}
}

// leakGoroutine 模拟创建大量 goroutine 导致的问题
func leakGoroutine() {
	time.Sleep(2 * time.Hour)
}
