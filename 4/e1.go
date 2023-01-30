// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
//которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора
//количества воркеров при старте.
//
//Программа должна завершаться по нажатию Ctrl+C.
//Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type WorkersPool struct {
	jobs chan int
}

// Начало работы Worker с чтением из канала и выводом в stdout

func (w *WorkersPool) WorkerStart(id int) {
	go func() {
		for work := range w.jobs {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(fmt.Sprintf("Worker %d: %d", id, work))
		}
	}()
}

func gen(c chan int) {
	c <- rand.Int()
}

func main() {
	var n int
	fmt.Print("Number of workers: ")
	fmt.Scanln(&n)

	wp := WorkersPool{
		jobs: make(chan int),
	}

	// Запись сигналов ОС
	osSig := make(chan os.Signal, 1)

	// Запись сигнала в канал
	signal.Notify(osSig, syscall.SIGINT)

	for i := 0; i < n; i++ {
		wp.WorkerStart(i + 1)
	}

loop:
	for { // При поступлении сигнала в канал обрабатывается завершение программы
		select {
		case <-osSig:
			{
				close(wp.jobs)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Shutting down...")
				break loop
			}
		default:
			gen(wp.jobs)
		}
	}

}
