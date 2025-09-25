package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep,
которая приостанавливает выполнение текущей горутины.

Важно: в отличии от настоящей time.Sleep, ваша функция должна именно блокировать выполнение
(например, через таймер или цикл), а не просто вызывать time.Sleep :) — это упражнение.

Можно использовать канал + горутину, или цикл на проверку времени (не лучший способ, но для обучения).
*/

func Sleep(milliseconds int) {
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)
		close(done)
	}()
	<-done
}

func main() {
	fmt.Println("Начинаю спать...")
	Sleep(2000) // (2 секунды)
	fmt.Println("Просыпаюсь!")
}
