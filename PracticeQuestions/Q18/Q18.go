package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

func main() {
	c := &counter{}
	wg := sync.WaitGroup{}

	// Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			c.incr()
			defer wg.Done()
		}()
	}

	// Channel
	c2 := NewChannelCounter()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			c2.Add(3)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(c.c)
	fmt.Println(c2.Read())
}

type counter struct {
	c  int
	mu sync.Mutex
	ch chan func()
}

func (c *counter) incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c++
}

func (c *counter) AddSome(input int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c += input
}

func NewChannelCounter() *counter {
	cntr := &counter{
		ch: make(chan func(), 100),
		c:  0,
	}
	go func(cntr *counter) {
		for f := range cntr.ch {
			f()
		}
	}(cntr)
	return cntr
}

func (c *counter) Add(input int) {
	c.ch <- func() {
		c.c = c.c + input
	}
}

func (c *counter) Read() int {
	ret := make(chan int)
	c.ch <- func() {
		ret <- c.c
		close(ret)
	}
	return <-ret
}
