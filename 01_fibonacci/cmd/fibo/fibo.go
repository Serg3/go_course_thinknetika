package main

import (
	"fmt"
	"go_course_thinknetika/01_fibonacci/pkg/fibo"
)

func main() {
	nums := []int{1, 3, 5, 8, 20, 22}
	for _, n := range nums {
		if n > 20 {
			fmt.Printf("Max number can be 20! You tried to reach %d number in sequence\n", n)
			continue
		}
		fmt.Printf("%d: %d\n", n, fibo.Seq(n))
	}
}
