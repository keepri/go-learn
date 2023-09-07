package mutex

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	mux   sync.RWMutex
	count int
}

func (c *counter) increment(n int) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.count += n
}

func (c *counter) peek() {
	c.mux.RLock()
	defer c.mux.RUnlock()
	fmt.Printf("count is %d\n", c.count)
}

func Run() {
	c := counter{count: 0}
	go c.increment(5)
    go c.peek()
	go c.increment(5)
	go c.peek()
	go c.increment(5)
	go c.increment(5)
	time.Sleep(time.Millisecond * 150)
    c.peek()
}
