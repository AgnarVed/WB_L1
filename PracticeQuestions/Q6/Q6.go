package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	stopChan := make(chan struct{})
	stoppedChan := make(chan struct{})
	//go func() {
	//	defer close(stoppedChan)
	//	fmt.Println("Starting work...")
	//	for {
	//		select {
	//		case <-stopChan:
	//			// stops
	//			return
	//		default:
	//			// TODO work
	//			fmt.Println("working...")
	//			time.Sleep(300 * time.Millisecond)
	//		}
	//	}
	//}()

	go go1(stopChan, stoppedChan, &wg)
	time.Sleep(2 * time.Second)
	log.Println("stopping goroutine 1...")

	close(stopChan) // tell it to stop
	<-stoppedChan   // wait for it to have stopped
	log.Println("Goroutine 1 stopped.")

	ctx, cancel := context.WithCancel(context.Background())

	go go2(ctx, &wg)
	time.Sleep(2 * time.Second)
	log.Println("stopping goroutine 2...")
	cancel()
	log.Println("Goroutine 2 stopped.")
	wg.Wait()

}

// Остановка с помощью канала
func go1(stopChan, stoppedChan chan struct{}, wg *sync.WaitGroup) {
	fmt.Println("Starting goroutine 1")
	defer close(stoppedChan)
	for {
		select {
		case <-stopChan:
			// stops
			wg.Done()
			return
		default:
			fmt.Println("Goroutine 1 is working...")
			time.Sleep(300 * time.Millisecond)
		}
	}

}

// Остановка с помощью контекста
func go2(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("Starting goroutine 2")
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			fmt.Println("Goroutine 2 is working...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}
