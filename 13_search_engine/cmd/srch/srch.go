package main

import (
	"go_course_thinknetika/13_search_engine/pkg/netsrv"
	"log"
)

func main() {
	l, err := netsrv.Listener("tcp4", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	sp := netsrv.New([]string{"https://golang.org", "https://go.dev"}, 2)
	netsrv.Searcher(l, sp)
}
