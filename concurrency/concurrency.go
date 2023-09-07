package concurrency

import (
	"fmt"
	"time"
)

func send_email(msg string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", msg)
	}()
	fmt.Printf("Email sent: '%s'\n", msg)
}

func Run() {
	send_email("Hello, world!")

	// added this so that the program doesn't quit
	// bafore the second print gets fired
	time.Sleep(time.Millisecond * 255)
}
