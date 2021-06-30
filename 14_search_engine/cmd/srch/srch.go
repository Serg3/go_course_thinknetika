package main

import (
	"go_course_thinknetika/14_search_engine/pkg/crawler"
	"go_course_thinknetika/14_search_engine/pkg/crawler/spider"
	"go_course_thinknetika/14_search_engine/pkg/webapp"
	"log"
)

func main() {
	docs := []crawler.Document{}
	go func() {
		docs = scan([]string{"https://golang.org", "https://go.dev"}, 2)
	}()

	paths := map[string]bool{
		"docs":  true,
		"index": true,
	}
	err := webapp.ListenAndServe(":8000", &docs, paths)
	if err != nil {
		log.Fatal(err)
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
