package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func sbor_with_chan(c chan int) string {
	summ := 0
	for val := range c {
		summ += val
	}
	return strconv.Itoa(summ)
}
func greet(mas []int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	summ := 0
	for _, i := range mas {
		summ += i * i

	}
	c <- summ

}
func main() {
	runtime.GOMAXPROCS(1)
	c := make(chan int, 2)
	mas := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(2)
	go greet(mas[len(mas)/2:], c, &wg)
	go greet(mas[:len(mas)/2], c, &wg)
	go func() {
		wg.Wait()
		close(c)
	}()

	fmt.Println(sbor_with_chan(c))

}
