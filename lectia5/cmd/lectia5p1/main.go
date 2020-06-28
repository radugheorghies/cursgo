package main

import (
	"log"
	"sync"
	"time"
)

var m sync.Mutex
var s int

var wg sync.WaitGroup

func main() {
	// var s int
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go sum(i)
	}

	// ch := make(chan int)

	// ch <- 10

	// v := <-ch

	wg.Wait()
	log.Println("S=", s)

}

func sum(i int) {
	m.Lock()
	s += i
	m.Unlock()
	wg.Done()
}

func showI(i int) {
	time.Sleep(time.Duration(i) * time.Second)
	log.Println(i)
}
