package main

import (
	"fmt"
	"strconv"
)

func vvod(c chan bool, chanel chan int) {
	for {
		select {
		case <-c:

			fmt.Println("Программа завершена!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			return
		case dann := <-chanel:

			fmt.Println(strconv.Itoa(dann) + "     " + strconv.Itoa(dann*dann))

		}
	}
}
func main() {
	c := make(chan bool)

	var numDann int
	chanel := make(chan int)
	flag := true
	go vvod(c, chanel)
	fmt.Println("Введите число: ")
	for flag {

		fmt.Scan(&numDann)
		if numDann == 0 {
			flag = false
			close(chanel)
			c <- true
			continue
		}
		chanel <- numDann

	}
}
