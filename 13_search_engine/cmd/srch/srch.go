package main

import (
	"go_course_thinknetika/13_search_engine/pkg/netsrv"
	"log"
)

func main() {
	err := netsrv.ListenAndSearch([]string{"https://golang.org", "https://go.dev"}, 2)
	// errors handler from net service
	if err != nil {
		log.Fatal(err)
	}
}
