// Реализовать структуру-счетчик, которая будет инкрементироваться
//в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	sync.Mutex
	cnt int
}

func (c *Counter) Inc() {
	c.Lock()
	c.cnt++
	c.Unlock()
}

type Counter2 struct {
	cnt int64
}

func (c *Counter2) Inc() {
	atomic.AddInt64(&c.cnt, 1)
}

func main() {
	n := 10000
	c1 := Counter{}
	c2 := Counter2{}
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			c1.Inc() // mutex
			c2.Inc() // atomic.AddInt64
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c1.cnt)
	fmt.Println(c2.cnt)
}
