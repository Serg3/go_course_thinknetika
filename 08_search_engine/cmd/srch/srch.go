package main

import (
	"flag"
	"fmt"
	"go_course_thinknetika/08_search_engine/pkg/crawler/spider"
	"go_course_thinknetika/08_search_engine/pkg/index"
	"log"
)

// Struct 'search' is used
// for a more convenient representation
// of search's parameters
type search struct {
	scanner *spider.Service
	sites   []string
	depth   int
	storage *index.Storage
}

func main() {
	param := flag.String("s", "", "word for search")
	flag.Parse()
	if *param == "" {
		flag.PrintDefaults()
		return
	}

	s := new()
	if s.storage.Empty() {
		fmt.Printf("Storage is empty. Performing a new scan...\n\n")
		for _, url := range s.sites {
			res, err := s.scanner.Scan(url, s.depth)
			if err != nil {
				log.Println(err)
				continue
			}
			s.storage.Append(res)
		}
		s.storage.Save()
	}
	s.storage.Index()
	s.storage.Sort()

	docs := s.storage.Search(param)
	fmt.Printf("Search results:\n\n")
	for _, d := range docs {
		fmt.Println(d)
	}
}

func new() *search {
	s := search{}
	s.sites = []string{"https://go.dev", "https://golang.org/"}
	s.depth = 2
	s.scanner = spider.New()
	s.storage = index.New()
	return &s
}
