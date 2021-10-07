package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // -1 tira um do contador
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1 tira um do contador
	}()

	go func() {
		escrever("Goroutine 3!")
		waitGroup.Done() // -1 tira um do contador
	}()

	go func() {
		escrever("Goroutine 4!")
		waitGroup.Done() // -1 tira um do contador
	}()

	waitGroup.Wait() // esperar a contagem chegar a zero
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
