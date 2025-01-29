package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := make(map[int]int)
	var mu sync.Mutex
	go func() {
		mu.Lock()
		for i := 0; i < 1000; i++ {

			m[i] = i * i

		}
		mu.Unlock()
	}()
	time.Sleep(1) //тут можно было бы поставить вэйт группу, потому что без слипа и без вейт группы планировщик решает сначала запустить основную горутину, ибо она бымтрее, но слип, в данном, случае, проще решает эту проблему
	mu.Lock()
	for key, value := range m {
		fmt.Println(key, value)
	}
	mu.Unlock()
	fmt.Println(len(m))
}
