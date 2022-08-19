package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	n := 5
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if len(os.Args) == 2 {
		n, err = strconv.Atoi(os.Args[1])
		if err != nil || n == 0 {
			fmt.Println("You need to write a number starting from 1")
			return
		}
	} else {
		fmt.Println("You need to specify 1 argument - amount of workers")
		return
	}

	workers := NewWorkers(n, cancel)
	workers.work(ctx)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		num := generator.Intn(1000)
		chars := 'a' + rune(generator.Intn('z'-'a'+1))
		select {
		case s := <-signals: // при получении сигнала завершаем работу всех воркеров
			fmt.Printf("\nProgramm has been stopped by signal %d\n", s)
			workers.stopWork()

			return
		default: // посылаем данные постоянно
			workers.sendData(num)
			workers.sendData(string(chars))
		}
		time.Sleep(time.Millisecond * 100)
	}

}

type Workers struct {
	ch     chan interface{}   // канал
	wg     sync.WaitGroup     // wg
	cancel context.CancelFunc // отмена
	n      int                // кол-во
}

func NewWorkers(n int, ctx context.CancelFunc) *Workers {
	return &Workers{
		ch:     make(chan interface{}, n),
		cancel: ctx,
		n:      n,
	}
}

func (w *Workers) work(ctx context.Context) {
	for i := 0; i < w.n; i++ {
		w.wg.Add(1)
		go func(id int) {
			for {
				select {
				// из канала пишем в data
				case data := <-w.ch:
					fmt.Printf("Worker %d read data from channel: Type %T, Value %v\n", id+1, data, data)
				// при получении Done пишем о завершении
				case <-ctx.Done():
					fmt.Println("Worker ", id+1, " stopped")
					return
				}
			}
			w.wg.Done()
		}(i)
	}
}

func (w *Workers) stopWork() {
	//w.cancel()  // ctx отмены
	w.wg.Wait() // wg
	close(w.ch) // закрываем канал
	fmt.Println("All workers have been stopped")
}

func (w *Workers) sendData(data interface{}) {
	w.ch <- data // отправляем данные в канал
}
