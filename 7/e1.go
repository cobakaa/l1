// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

type SyncMap struct {
	sync.RWMutex // Встраивание Mutex
	data         map[int]int
}

func (sm *SyncMap) Set(k int, v int) {
	sm.Lock()
	sm.data[k] = v
	sm.Unlock()
}

func (sm *SyncMap) Add(k int, v int) {
	sm.Lock()
	sm.data[k] += v
	sm.Unlock()
}

func main() {
	tmp := make(map[int]int)

	sm := SyncMap{data: tmp}

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(k int) {
			sm.Set(k%10, k)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(sm.data)

	// Новые значения всегда на 2 больше предыдущих
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(k int) {
			sm.Add(k%10, 1)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(sm.data)

}
