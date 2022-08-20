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

var justString string

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

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
	v := createHugeString(1 << 10)

	last := 100
	switch {
	case last <= len(v)-1:
		justString = v[:100]
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
	fmt.Println("cap(v) ", cap([]rune(v)))
	fmt.Printf("%p\n", &v)
	fmt.Println("Size in bytes ", int(unsafe.Sizeof(v))+len(v))
}

func main() {

	str := "avsdawefasdfdsfaec"
	nstr := str[:5]
	fmt.Println(nstr)
	fmt.Println(str)
	fmt.Println(len(nstr))
	fmt.Println(&nstr)
	fmt.Println(&str)
	//someFunc()
	//fmt.Println(justString)
	//fmt.Println(len(justString))
	//fmt.Println(createHugeString(1 << 20))
}
