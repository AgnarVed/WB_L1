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
	var n int     // кол-во воркеров
	var err error // ошибка

	ctx, cancel := signal.NotifyContext(context.Background(), os.Signal(syscall.SIGINT))
	defer cancel()

	// проверка аргументов
	if len(os.Args) == 2 {
		n, err = strconv.Atoi(os.Args[1])
		if err != nil || n < 0 {
			fmt.Println("Введите число от 1")
			return
		}
	} else {
		fmt.Println("Уточните кол-во воркеров")
		return
	}

	workersCh := make(chan interface{}, n)

	var wg sync.WaitGroup
	work(ctx, workersCh, n, &wg)

	go func() {
		generator := rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			num := generator.Intn(100)
			//chars := 'a' + rune(generator.Intn('z'-'a'+1))
			select {
			case <-ctx.Done(): // при получении сигнала завершаем работу всех воркеров
				fmt.Printf("Programm has been stopped by signal\n")
				close(workersCh)
				return
			default: // посылаем данные постоянно
				workersCh <- num
				//workersCh <- string(chars)
			}
			time.Sleep(time.Millisecond * 300)
		}
	}()

	wg.Wait() // wg
	// закрываем канал
	fmt.Println("\nAll workers have been stopped")
}

func work(ctx context.Context, ch <-chan interface{}, n int, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				// из канала пишем в data
				case data := <-ch:
					fmt.Printf("Worker %d обработал: Type %T, Value %v\n", id+1, data, data)
					// при получении Done
				case <-ctx.Done():
					return
				}
			}
		}(i)
	}
}
