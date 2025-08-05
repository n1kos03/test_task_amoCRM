package main

import (
	"fmt"
)

// Merge принимает два read-only канала и возвращает выходной канал,
// в который последовательно (в любом порядке) будут отправлены все значения
// из обоих входных каналов.
//
// Выходной канал должен быть закрыт после того, как оба входных канала закроются.
// Merge не должен закрывать входные каналы
//
// Для проверки решения запустите тесты: go test -v
func Merge(ch1, ch2 <-chan int) <-chan int {
	// TODO: реализуйте эту функцию
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for {
			v1, ok1 := <-ch1
			v2, ok2 := <-ch2
			if ok1 {
				outCh <- v1
			}
			if ok2 {
				outCh <- v2
			}
			if !ok1 && !ok2 {
				break
			}
		}
	}()
	return outCh
}

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		defer close(a)
		a <- 4
		a <- 1
	}()

	go func() {
		defer close(b)
		b <- 2
		b <- 4
	}()

	for v := range Merge(a, b) {
		fmt.Println(v)
	}
}
