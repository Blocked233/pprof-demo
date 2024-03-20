package main

import (
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	for {
		calculatePi()
		time.Sleep(100 * time.Millisecond)
	}
}

// calculatePi 模拟一个 CPU 密集型的函数
func calculatePi() float64 {
	var total float64
	for i := 0; i < 1000000; i++ {
		total += rand.Float64()
	}
	return total
}
