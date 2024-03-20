package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	for {
		leakMemory()
	}
}

// leakMemory 模拟一个内存分配问题
func leakMemory() {
	data := make([]byte, 1000000)
	_ = data
}
