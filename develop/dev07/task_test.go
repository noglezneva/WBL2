package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	// Создаем каналы
	ch1 := sig(2 * time.Hour)
	ch2 := sig(5 * time.Minute)
	ch3 := sig(1 * time.Second)
	ch4 := sig(1 * time.Hour)
	ch5 := sig(1 * time.Minute)

	// Объединяем каналы с помощью функции or
	out := or(ch1, ch2, ch3, ch4, ch5)

	// Ждем, пока закроется объединенный канал
	<-out

	elapsed := time.Since(start)

	// Проверяем, что прошло не менее 5 минут и не более 2 часов
	if elapsed < 5*time.Minute || elapsed > 2*time.Hour {
		t.Errorf("Unexpected time elapsed: %v", elapsed)
	}
}
