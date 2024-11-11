package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	for num, val := range []int{2, 4, 6, 8, 10} {
		go fmt.Printf("%d. Значение квадрата %d равно %d\n", num, val, val*val)

	}
	time.Sleep(100000)
}
