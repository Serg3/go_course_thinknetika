package main

import (
	"fmt"
	"sync"
)

func game(wg *sync.WaitGroup, p string) {
	fmt.Println("Player", p)
	wg.Done()
}

func main() {
	p1, p2 := "Biba", "Boba"

	var wg sync.WaitGroup
	wg.Add(2)
	go game(&wg, p1)
	go game(&wg, p2)
	wg.Wait()
}
