package index

import (
	"fmt"
	"go_course_thinknetika/04_search_engine/pkg/crawler"
	"hash/fnv"
	"sort"
	"strings"
)

type crwDocs []crawler.Document

// Struct 'Storage' is used
// for a more convenient representation
// of storages's parameters
type Storage struct {
	counter int
	docs    crwDocs
	ind     map[uint32][]int
}

// New creates a new storage instance
func New() *Storage {
	return &Storage{
		counter: 0,
		docs:    make(crwDocs, 0),
		ind:     make(map[uint32][]int),
	}
}

// Append adds document to the storage
func (s *Storage) Append(docs []crawler.Document) {
	for _, d := range docs {
		s.counter++
		d.ID = s.counter
		s.docs = append(s.docs, d)
		s.index(d.ID, d.Title)
	}
}

// Search performs a search
// by the incoming param
// in the indexed storage
// and returns formatted result
func (s *Storage) Search(param *string) []string {
	var d crawler.Document

	res := make([]string, 0)
	h := hash(strings.ToLower(*param))
	ids := s.ind[h]

	for _, id := range ids {
		d = s.binarySearch(id, 0, len(s.docs))
		if d.ID != 0 {
			res = append(res, fmt.Sprintf("%d: %s (%s)", d.ID, d.Title, d.URL))
		}
	}

	return res
}

func (s *Storage) Sort() {
	sort.Sort(s.docs)
}

func (s *Storage) binarySearch(id, l, r int) crawler.Document {
	if r < l {
		return crawler.Document{}
	}

	mid := l + (r-l)/2

	if s.docs[mid].ID == id {
		return s.docs[mid]
	}

	if id <= s.docs[mid].ID {
		return s.binarySearch(id, l, mid-1)
	} else {
		return s.binarySearch(id, mid+1, r)
	}
}

func (s *Storage) index(id int, title string) {
	var h uint32

	arr := strings.Split(title, " ")
	for _, t := range arr {
		h = hash(strings.ToLower(t))
		if h > 0 {
			if intArr, ok := s.ind[h]; !ok {
				intArr = make([]int, 0)
				intArr = append(intArr, id)
				s.ind[h] = intArr
			} else {
				intArr = append(intArr, id)
				s.ind[h] = intArr
			}
		}
	}
}

// Hash generation has been taken
// from the stackoverflow
// with help of "hash/fnv" package
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// Below methods are needed
// for using sort.Interface
func (d crwDocs) Len() int           { return len(d) }
func (d crwDocs) Less(i, j int) bool { return d[i].ID < d[j].ID }
func (d crwDocs) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
