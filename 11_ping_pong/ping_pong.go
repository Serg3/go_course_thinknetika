package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const maxScore = 11

func player(c chan string, wg *sync.WaitGroup, name string) {
	for {
		s := <-c

		if s == "stop" {
			wg.Done()
			c <- name
			return
		}
		if s != "begin" && rand.Intn(100)%5 == 0 {
			c <- "stop"
			wg.Done()
			return
		}
		if s == "ping" {
			s = "pong"
		} else {
			s = "ping"
		}

		c <- s
		fmt.Printf("%s: %s!\n", name, s)
	}
}

func main() {
	p1, p2 := "Biba", "Boba"
	var winner string
	res := map[string]int{
		p1: 0,
		p2: 0,
	}
	game := make(chan string)
	wg := sync.WaitGroup{}

	for res[p1] < maxScore && res[p2] < maxScore {
		fmt.Println("-= Begin! =-")

		wg.Add(2)
		go player(game, &wg, p1)
		go player(game, &wg, p2)
		game <- "begin"
		wg.Wait()

		loser := <-game
		if loser == p1 {
			winner = p2
		} else {
			winner = p1
		}
		res[winner]++
		fmt.Printf("-= Goal! %s +1 =-\n", winner)
		fmt.Printf("%s %d : %d %s\n\n", p1, res[p1], res[p2], p2)
	}

	fmt.Printf("-= Match is over! %s wins! =-\n", winner)
}
