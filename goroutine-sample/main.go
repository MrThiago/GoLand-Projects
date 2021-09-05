package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Start Process")

	go SampleOne()
	go SampleOne()
	go SampleOne()

	//SampleTwo()

	fmt.Println("End Process")
	fmt.Scanln()
}

func BackgroundOperation(i int) {
	time.Sleep(time.Second)
	fmt.Println("Background Operation Done", i)
}

func SampleOne() {
	for i := 0; i < 10; i++ {
		BackgroundOperation(i)
	}
}

func SampleTwo() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			BackgroundOperation(i)
			wg.Done()
		}()
	}

	wg.Wait()
}
