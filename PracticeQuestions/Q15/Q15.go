package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"unsafe"
)

//
// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации
// var justString string
// func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
// }
//
// func main() {
// someFunc()
// }

// может возникнуть проблема взятия строки ( если текст не UTF-8, то по байтам бука может не поместиться, отсюда неверно взятое значение)

var justString string
var justString2 string

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZπ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// createHugeString создает большую строку с размером size (можно было и обычным способом сгенерировать,
// нашел на stackoverflow этот вариант
func createHugeString(size int) string {
	b := make([]byte, size)
	for i, cache, remain := size-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

func someFunc() {
	//v := createHugeString(1 << 10)
	v := "abcdeπfghijkπlmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZπ"
	runes := []rune(v)
	last := 25
	switch {
	case last < len(v):
		justString = string(runes[:10])
	default:
		log.Println("Выход за пределы строки")
	}

	fmt.Println("justString = ", justString)
	fmt.Println("len(justString) = ", len(justString))
	fmt.Println("cap(justString) ", cap([]rune(justString)))
	fmt.Printf("%p\n", &justString)
	fmt.Println("Size in bytes ", int(unsafe.Sizeof(justString))+len(justString))
	fmt.Println()
	fmt.Println("v = ", v)
	fmt.Println("len(v) = ", len(v))
	fmt.Println("cap(v) ", cap([]rune(v)))
	fmt.Printf("%p\n", &v)
	fmt.Println("Size in bytes ", int(unsafe.Sizeof(v))+len(v))
}

func main() {

	str := "apple_π!"
	subStr := str[0:7]
	fmt.Println(subStr)

	for i, val := range str {
		fmt.Printf("index: %d value: %d char %s\n", i, val, string(val))
	}
	someFunc()
}
