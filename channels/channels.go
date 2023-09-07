package channels

import (
	"fmt"
	"time"
)

func Run() {
	ch := make(chan int)
	go func(ch chan int) {
		ch <- 42
		time.Sleep(time.Millisecond * 250)
		ch <- 69
	}(ch)

	// this will wait until it has something to dequeue
	fmt.Printf("val one %d\n", <-ch)
	// this will wait until it has something to dequeue
	fmt.Printf("waiting for val two to be set...\n")
	fmt.Printf("val two %d\n", <-ch)

	num_dbs := 5
	db_chan := get_databases_channel(&num_dbs)
	wait_dbs(&num_dbs, db_chan)

	concurrent_fib(10)

	ch_one := make(chan string)
	ch_two := make(chan string)
	go listen_channels(&ch_one, &ch_two)
	ch_two <- "two"
	ch_two <- "two"
	ch_two <- "two"
	ch_one <- "one"
	close(ch_one)
	close(ch_two)

	// keep in scope
	time.Sleep(time.Millisecond * 50)
}

func get_databases_channel(num_dbs *int) *chan struct{} {
	ch := make(chan struct{})
	// or
	// ch := make(chan struct{}, *num_dbs)
	// using buffered channels
	// allows us to stay on the same thread
	go connect_dbs(num_dbs, &ch)
	return &ch
}

func wait_dbs(num_dbs *int, ch *chan struct{}) {
	for i := 0; i < *num_dbs; i++ {
		<-*ch
	}
}

func connect_dbs(num_dbs *int, ch *chan struct{}) {
	for i := 0; i < *num_dbs; i++ {
		*ch <- struct{}{}
		fmt.Printf("Database %d is online\n", i+1)
		time.Sleep(time.Millisecond * 7)
	}
	close(*ch)
}

func concurrent_fib(n int) {
	ch := make(chan int, n)
	go fibonacci(n, &ch)
	for v := range ch {
		fmt.Printf("fib %d\n", v)
	}
}

func fibonacci(n int, ch *chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		*ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 7)
	}
	close(*ch)
}

func listen_channels(ch_one, ch_two *chan string) {
	select {
	case v, ok := <-*ch_one:
		if !ok {
			fmt.Printf("could not read channel one")
		}
		fmt.Printf("channel %v triggered\n", v)
		return

	case v, ok := <-*ch_two:
		if !ok {
			fmt.Printf("could not read channel two")
		}
		fmt.Printf("channel %v triggered\n", v)
		go listen_channels(ch_one, ch_two)
	}

}
