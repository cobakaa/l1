package main

import (
	"fmt"
	"time"
)

func main() {
	f := false

	// Горутана завершает работу при условии изменения значения переменной
	go func() {
		for {
			if f {
				return
			}

			time.Sleep(200 * time.Millisecond)
			fmt.Println("Working...")
		}
	}()

	time.Sleep(800 * time.Millisecond)
	f = true

}
