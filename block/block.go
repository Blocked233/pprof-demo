package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

var (
	myLock sync.Mutex
)

// lockOperation 模拟一个会阻塞的操作
func lockOperation() {
	myLock.Lock()
	defer myLock.Unlock()

	// 模拟长时间操作，以引起阻塞
	time.Sleep(5 * time.Second)
	fmt.Println("Operation completed")
}

func main() {
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪，block
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪，mutex

	// 启动 HTTP 服务器以供 pprof 使用
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	// 并发执行 lockOperation 来模拟阻塞
	for i := 0; i < 1000; i++ {
		go lockOperation()
	}

	// 保持程序运行，以便进行 pprof 分析
	select {}
}
