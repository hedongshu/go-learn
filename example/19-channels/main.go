package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	msg := make(chan string)

	go func() {
		msg <- "ping"
	}()

	req := <-msg
	println(req)

	// 通道缓存
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// 通道同步
	done := make(chan bool, 1)
	go worker(done)

	status := <-done
	fmt.Println("status:", status)

	// 通道方向
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)

}
