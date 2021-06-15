package main

import (
	"flag"
	"fmt"
	"go_course_thinknetika/04_search_engine/pkg/crawler/spider"
	"go_course_thinknetika/04_search_engine/pkg/index"
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
	if *param != "" {
		docs := scan().Search(param)
		fmt.Println("Search results:")
		for _, d := range docs {
			fmt.Println(d)
		}
	} else {
		flag.PrintDefaults()
	}
}

// Function scan() uses package 'crawler'
// to search through Go sites by word
// and returs sorted *index.Storage result
func scan() *index.Storage {
	s := new()
	for _, url := range s.sites {
		res, err := s.scanner.Scan(url, s.depth)
		if err != nil {
			log.Println(err)
			continue
		}
		s.storage.Append(res)
	}
	s.storage.Sort()
	return s.storage
}

func new() *search {
	s := search{}
	s.sites = []string{"https://go.dev", "https://golang.org/"}
	s.depth = 2
	s.scanner = spider.New()
	s.storage = index.New()
	return &s
}
