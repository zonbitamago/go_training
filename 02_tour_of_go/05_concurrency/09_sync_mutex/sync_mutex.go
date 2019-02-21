package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 排他制御カウンタ
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 排他制御カウンタ増加処理
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// Value 排他制御カウンタ値取得処理
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
