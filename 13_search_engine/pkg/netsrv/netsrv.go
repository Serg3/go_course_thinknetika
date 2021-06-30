package netsrv

import (
	"bufio"
	"go_course_thinknetika/13_search_engine/pkg/crawler"
	"net"
	"strings"
)

// ListenAndSearch accepts two parameters for network (string) and address (string).
// Returns error if something went wrong with network connection.
func ListenAndSearch(network, address string) (net.Listener, error) {
	l, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Function Searcher() performs handling of incoming connections
// and makes a search by received word through provided documents.
func Searcher(conn net.Conn, docs []crawler.Document) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		conn.Write([]byte("\nEnter a word: "))
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		if len(msg) > 0 {
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
