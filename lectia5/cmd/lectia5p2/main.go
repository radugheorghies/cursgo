package main

import (
	"log"
	"time"
)

type CustomChan struct {
	Text         string
	ResponseChan chan string
}

func main() {

	cChan := make(chan CustomChan)

	go func() {
		for v := range cChan {
			log.Println(v.Text)
			time.Sleep(time.Millisecond * 500)
			v.ResponseChan <- "ok then."
		}
	}()

	go func() {
		for {
			localChan := make(chan string)

			myObj := CustomChan{
				Text:         "hi there",
				ResponseChan: localChan,
			}

			cChan <- myObj

			for v := range localChan {
				log.Println(v)
				close(localChan)
			}

			//time.Sleep(time.Millisecond * 500)
		}
	}()

	select {}

	// cPing := make(chan string)
	// cPong := make(chan string)

	// for i := 0; i < 10; i++ {
	// 	go func(nr int) {
	// 		for v := range cPing {
	// 			log.Println("cPing chan from goroutine ", nr, ":", v)
	// 			cPong <- "pong"
	// 		}
	// 	}(i)
	// }

	// go func() {
	// 	for {
	// 		cPing <- "ping"
	// 		time.Sleep(time.Millisecond * 500)
	// 	}
	// }()

	// for v := range cPong {
	// 	log.Println("Pong from main func:", v)
	// }

}
