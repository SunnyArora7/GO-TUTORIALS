package main

import (
	"fmt"
	"sync"
)

// We use mutex suppose 1000 of persons hit a like button at the same time to the increment value , and we try to get the value it will display everytime different so get the consistent value we use mutex, only lock that much part ie need to be consistent
type post struct {
	likes int
	mu    sync.Mutex
}

func (p *post) incLikes(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.likes++

}
func main() {
	var wg sync.WaitGroup
	post1 := post{likes: 0}
	for range 100 {
		wg.Add(1)
		go post1.incLikes(&wg)

	}
	wg.Wait()
	fmt.Println(post1.likes)
}
