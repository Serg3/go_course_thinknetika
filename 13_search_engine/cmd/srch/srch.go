package main

import (
	"go_course_thinknetika/13_search_engine/pkg/crawler"
	"go_course_thinknetika/13_search_engine/pkg/crawler/spider"
	"go_course_thinknetika/13_search_engine/pkg/netsrv"
	"log"
	"sync"
)

func main() {
	var docs []crawler.Document
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		docs = scan([]string{"https://golang.org", "https://go.dev"}, 2)
		wg.Done()
	}()

	listener, err := netsrv.ListenAndSearch("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		if len(docs) == 0 {
			conn.Write([]byte("Buffer is empty. A new scan is in progress...\n"))
			wg.Wait()
		}
		go netsrv.Searcher(conn, docs)
	}
}

func scan(urls []string, depth int) (docs []crawler.Document) {
	scn := spider.New()
	for _, url := range urls {
		res, err := scn.Scan(url, depth)
		if err != nil {
			log.Println(err)
			continue
		}
		docs = append(docs, res...)
	}
	return docs
}
