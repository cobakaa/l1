// Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	workingTime := 4

	ch := make(chan int)

	// Отправление в канал
	go func() {
		for {
			ch <- rand.Int()
		}
	}()

	// Считывание из канала
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	time.Sleep(time.Duration(workingTime) * time.Second)

}
