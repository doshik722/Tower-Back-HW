package main

import (
	"fmt"
	"runtime"
	"time"
)

func Besk(j int) {
	if j > 99 {
		runtime.Goexit()

	}
	fmt.Println(j)
	j++
	Besk(j)
}
func main() {
	/* завершение через дополнительный канал и конструкцию select
	quit := make(chan struct{})
	i := 0
	go func() {
		for {
			select {
			case <-quit:
				return
			default:

				for {
					fmt.Println(i)
					i++
				}
			}
		}
	}()
	time.Sleep(1)
	close(quit)
	*/
	/*
		j := 0 // реализация функции Besk выше, но суть в том, что я заканчиваю ее с ппомощью runtime.Goexit()
		go Besk(j)

		time.Sleep(100000)
	*/
	var stop bool = true //через флаги

	go func() {
		i := 0
		for stop {
			fmt.Println(i)
			i++

		}
	}()
	time.Sleep(1)
	stop = false
}
