package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := make(chan bool)

	numDann := 0
	chanel := make(chan string)
	flag := true //сигнал, что в мэйне данные записывать в канал больше НЕЛЬЗЯ

	go func() {
		for {
			select {
			case <-c:
				dann, ok := <-chanel
				if ok {
					fmt.Println(strconv.Itoa(1) + " воркер: " + dann)

				}
				fmt.Println("Программа завершена!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!") //почему-то выводится только
				//через дебаггер, но не через терминал. Кроме того, дебаггер успевает выполнить на ~40 тыс операций больше, чем терминал. Почему?
				return
			case dann := <-chanel:

				fmt.Println(strconv.Itoa(1) + " воркер: " + dann)

			}
		}
	}()

	go func() {
		time.Sleep(time.Second)
		flag = false
		c <- true

		close(chanel)
	}()
	for flag { //делаем произвольные данные и заносим их в канал
		chanel <- ("Произвольные данные " + strconv.Itoa(numDann))
		numDann++

	}

}
