package sync

import "sync"

// type Counter interface {
// 	Inc() int
// 	value() int
// }

type Counter struct {
	mu sync.Mutex
	// Alternatively, we can embed `sync.Mutex` here, i.e. don't give it a name.
	// This way we can just call Lock/Unlock on the Counter object.
	// However, this way the methods of the embeded lock becomes part of the public interface.
	value int
}

// This is what a constructor would look like in golang
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += 1
}

func (c *Counter) Value() int {
	return c.value
}

// # When to use locks over channels and goroutines?
// A common Go newbie mistake is to over-use channels and goroutines just because it's possible,
// and/or because it's fun. Don't be afraid to use a sync.Mutex if that fits your problem best.
// Go is pragmatic in letting you use the tools that solve your problem best and not forcing you
// into one style of code.
// Paraphrasing:
// - Use channels when passing ownership of data
// - Use mutexes for managing state
