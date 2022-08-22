package main

import (
	"sync"
	"testing"
)

func testCorrectness(t *testing.T, cntr counter) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		if i%3 == 0 {
			go func(counter counter) {
				counter.Read()
				wg.Done()
			}(cntr)
		} else if i%3 == 1 {
			go func(counter counter) {
				counter.Add(1)
				counter.Read()
				wg.Done()
			}(cntr)
		} else {
			go func(counter counter) {
				counter.Add(1)
				wg.Done()
			}(cntr)
		}
	}

	wg.Wait()

	if cntr.Read() != 66 {
		t.Errorf("counter should be %d and was %d", 66, cntr.Read())
	}
}
