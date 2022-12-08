package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ctrlc := make(chan os.Signal, 1)
	signal.Notify(ctrlc, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("Нажмите Ctrl+c для остановки")
	go func() {
		var x int
		for {
			select {
			case <-ctrlc:
				fmt.Println("")
				fmt.Println("Конец")
				return
			default:
				x = x + 1
				fmt.Println(x * x)
				time.Sleep(time.Second)
			}
		}
	}()
	wg.Wait()
}
