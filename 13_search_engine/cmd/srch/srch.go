package main

import (
	"go_course_thinknetika/13_search_engine/pkg/netsrv"
	"log"
)

func main() {
	// listener, err := net.Listen("tcp4", ":12345")
	l, err := netsrv.Listener("tcp4", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	netsrv.Searcher(l)
}
