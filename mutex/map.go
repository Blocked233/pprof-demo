package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
)

var (
	lock sync.Mutex
	data = make(map[int]int)
)

func main() {
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪，block
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪，mutex
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for i := 0; i < 1000; i++ {
		go worker(i)
	}

	select {}
}

func worker(id int) {
	for {
		lock.Lock()
		data[id] = id
		lock.Unlock()
	}
}
