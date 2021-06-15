package main

import (
	"flag"
	"fmt"
	"go_course_thinknetika/04_search_engine/pkg/crawler"
	"go_course_thinknetika/04_search_engine/pkg/crawler/spider"
	"log"
	"strings"
)

func main() {
	param := flag.String("s", "", "word for search")
	flag.Parse()
	if *param != "" {
		docs := scan()
		fmt.Println("Search result:")
		for _, doc := range docs {
			if strings.Contains(strings.ToLower(doc.Title), strings.ToLower(*param)) {
				fmt.Println(doc.URL, doc.Title)
			}
		}
	} else {
		flag.PrintDefaults()
	}
}

// Function scan() uses package 'crawler'
// to search through Go sites by word
// and returs []crawler.Document result
func scan() (docs []crawler.Document) {
	scn := spider.New()
	const depth = 2
	urls := []string{"https://golang.org", "https://go.dev"}
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
