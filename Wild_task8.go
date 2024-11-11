package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите число")
	var x int64
	fmt.Scan(&x)
	fmt.Println("Введите бит, который хотите заменить")
	var i int
	fmt.Scan(&i)
	fmt.Println("Введите заменy")
	var zam int
	fmt.Scan(&zam)
	if i < 0 || i > 63 {
		fmt.Println("Ошибка: позиция бита должна быть в диапазоне от 0 до 63.")
		return
	}
	binarymas := []rune(fmt.Sprintf("%064b", uint64(x)))
	binarymas[i] = rune(zam + '0')
	binaryStringNew := string(binarymas)
	fmt.Println(binaryStringNew)
}
