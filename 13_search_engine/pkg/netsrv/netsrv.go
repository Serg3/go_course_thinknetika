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

// ListenAndSearch accepts two parameters for URLs ([]string) and depth (int),
// performs a new scan of provided sites with specific depth
// and listen "tcp" network addresses on port ":8000".
// Returns error if something went wrong with network connection.
func ListenAndSearch(urls []string, depth int) error {
	sp := SearchParams{}
	sp.urls = urls
	sp.depth = depth

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		return err
	}

	err = handler(l, sp)
	if err != nil {
		return err
	}

	return nil
}

// Handler performs handling of incoming connections
// from the listener and parameters SearchParams
// given to the function
// and provides a search result from scanned sites
// by incoming word from the reader.
func handler(listener net.Listener, sp SearchParams) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go searcher(conn, sp)
	}
}

func searcher(conn net.Conn, sp SearchParams) {
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
