package main

import (
	"sync"
	"sync/atomic"
	"time"
)

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	arr := []int64{2, 4, 6, 8, 10}
	WithMutex(arr)  // конкурентно с мьютексом
	WithAtomic(arr) // конкурентно с атомик
	WithoutAny(arr) // неконкурентно
}

func (s *Summ) AddValue(input int64) {
	s.mu.Lock()
	s.value += input * input
	s.mu.Unlock()
}

func (s *Summ) GetValue() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.value
}

type Summ struct {
	value int64
	mu    sync.Mutex
}

func WithMutex(arr []int64) int64 {
	res := Summ{value: 0}
	for _, val := range arr {
		go res.AddValue(val)
	}
	time.Sleep(2 * time.Second)
	//fmt.Println(res.GetValue())
	return res.GetValue()
}

func WithAtomic(arr []int64) int64 {
	var res int64
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for _, val := range arr {

		go func(val int64) {
			inc := val * val
			atomic.AddInt64(&res, inc)
			wg.Done()
		}(val)
	}
	wg.Wait()
	//fmt.Println(res)
	return res
}

func WithoutAny(arr []int64) int64 {
	var res int64
	for i := 0; i < len(arr); i++ {
		val := arr[i]
		res += val * val
	}
	//fmt.Println(res)
	return res
}
