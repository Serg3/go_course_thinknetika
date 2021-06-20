package main

import (
	"flag"
	"fmt"
	"go_course_thinknetika/08_search_engine/pkg/crawler/spider"
	"go_course_thinknetika/08_search_engine/pkg/index"
	"go_course_thinknetika/08_search_engine/pkg/storage"
	"log"
	"os"
)

const file = "scan_result.txt"

// Struct 'search' is used
// for a more convenient representation
// of search's parameters
type search struct {
	scanner *spider.Service
	sites   []string
	depth   int
	index   *index.Store
	storage *storage.Filestore
}

func main() {
	param := flag.String("s", "", "word for search")
	flag.Parse()
	if *param == "" {
		flag.PrintDefaults()
		return
	}

	s := new()

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("File not found:", err.Error())
	} else {
		data, err := s.storage.Load(f)
		if err != nil {
			fmt.Println("Could not read a file:", err.Error())
		}
		s.index.Append(data)
	}
	defer f.Close()

	if s.index.Empty() {
		fmt.Printf("Storage is empty. Performing a new scan...\n\n")
		for _, url := range s.sites {
			res, err := s.scanner.Scan(url, s.depth)
			if err != nil {
				log.Println(err)
				continue
			}
			s.index.Append(res)
		}
	}
	s.index.Index()
	s.index.Sort()

	res := s.index.Search(param)
	fmt.Printf("Search results:\n\n")
	for _, d := range res {
		fmt.Println(d)
	}

	f, err = os.Create(file)
	if err != nil {
		fmt.Println("Could not create a file:", err.Error())
	}

	s.storage.Save(f, s.index.Docs())
}

func new() *search {
	s := search{}
	s.sites = []string{"https://go.dev", "https://golang.org/"}
	s.depth = 2
	s.scanner = spider.New()
	s.index = index.New()
	s.storage = storage.New()
	return &s
}
