package main

import (
	"log"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "hi"
	}()
	go func() {
		ch2 <- "hi 2"
	}()

	select {
	case v := <-ch1:
		log.Println("From ch1:", v)
	case v := <-ch2:
		log.Println("From ch2:", v)
	default:
		log.Println("NU AM PRIMIT NIMIC INTERESANT")
	}
}
