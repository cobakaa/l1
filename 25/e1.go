// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func sleep1(t time.Duration) {
	<-time.After(t) // same NewTimer(t).C
}

func sleep2(t time.Duration) {
	begin := time.Now()
	for time.Now().Sub(begin) < t { // проверка, сколько времени прошло с момента захода в функцию
	}
}

func main() {
	t := time.Now()
	sleep1(2 * time.Second)
	fmt.Println(time.Since(t))
	t = time.Now()
	sleep2(2 * time.Second)
	fmt.Println(time.Since(t))
}
