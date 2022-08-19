package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map

func main() {
	m := NewMap()
	arr := []int{1, 2, 3, 4, 5}
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for i := 0; i < len(arr); i++ {
		go m.addToMap(&wg, arr[i], arr[i]*2)
	}
	wg.Wait()

	fmt.Println(m.mp[5])
}

type myMap struct {
	mp map[int]int
	mu sync.Mutex
}

func NewMap() myMap {
	return myMap{
		mp: make(map[int]int),
		mu: sync.Mutex{},
	}
}

func (m *myMap) addToMap(wg *sync.WaitGroup, key, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mp[key] = value
	wg.Done()
}
