// Implement a thread-safe counter using Mutex
package main
import (
	"fmt"
	"sync"
)

type Counter struct { // Struct Counter chứa Mutex(mt) và biến đếm (value)
	mt sync.Mutex
	value int
}

func (c * Counter) Increase() { //Hàm cộng biến đếm
	c.mt.Lock()
	c.value++
	c.mt.Unlock()
}

// Hàm lấy giá trị hiện tại
func (c *Counter) Value() int {
	c.mt.Lock()
	defer c.mt.Unlock() 
	return c.value
}

func main() {
	counter := Counter{}
	var wait_group sync.WaitGroup
	for i:=0; i<10000; i++ {
		wait_group.Add(1)
		go func () {
			counter.Increase()
			wait_group.Done()
		}()
	}

	wait_group.Wait()
	fmt.Printf("Counter = %d\n", counter.Value())
}
