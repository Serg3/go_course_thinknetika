package netsrv

import (
	"bufio"
	"go_course_thinknetika/13_search_engine/pkg/crawler"
	"go_course_thinknetika/13_search_engine/pkg/crawler/spider"
	"log"
	"net"
	"strings"
)

type SearchParams struct {
	urls  []string
	depth int
}

func New(urls []string, depth int) SearchParams {
	sp := SearchParams{}
	sp.urls = urls
	sp.depth = depth
	return sp
}

// Listener accepts two parameters
// for address (string) and port (string)
// and returns listener of the local system.
func Listener(network, address string) (net.Listener, error) {
	return net.Listen(network, address)
}

// Searcher performs handling of incoming connections
// from the listener and parameters SearchParams
// given to the function
// and returns a search result from scanned sites
// by incoming word from the reader.
func Searcher(listener net.Listener, sp SearchParams) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handler(conn, sp)
	}
}

func handler(conn net.Conn, sp SearchParams) {
	var docs []crawler.Document
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		conn.Write([]byte("\nEnter a word: "))
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		if len(msg) > 0 {
			if len(docs) == 0 {
				conn.Write([]byte("Buffer is empty. Performing a new scan...\n"))
				docs = scan(sp)
			}
			conn.Write([]byte("Search result:\n"))
			for _, doc := range docs {
				if strings.Contains(strings.ToLower(doc.Title), strings.ToLower(string(msg))) {
					res := doc.URL + doc.Title + "\n"
					conn.Write([]byte(res))
				}
			}
		} else {
			// for easy exit
			conn.Write([]byte("Empty search!\n"))
			return
		}
	}
}

// Function scan() uses package 'crawler'
// to search through Go sites by word
// and returs []crawler.Document result
func scan(sp SearchParams) (docs []crawler.Document) {
	scn := spider.New()
	for _, url := range sp.urls {
		res, err := scn.Scan(url, sp.depth)
		if err != nil {
			log.Println(err)
			continue
		}
		docs = append(docs, res...)
	}
	return docs
}
