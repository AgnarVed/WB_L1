package main

import (
	"sync"
	"testing"
)

func BenchmarkSqr(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	for i := 0; i < b.N; i++ {
		Sqr(arr)
	}
}

func BenchmarkSqrConc2(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		SqrConc(arr, &wg)
	}
	wg.Wait()
}
