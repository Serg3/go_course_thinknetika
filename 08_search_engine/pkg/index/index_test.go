package index

import (
	"go_course_thinknetika/08_search_engine/pkg/crawler"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

// Tests

func TestStore_Docs(t *testing.T) {
	store := New()
	docs := crwDocs{{ID: 1}}
	store.Append(docs)

	got := store.Docs()
	want := docs
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestStore_Append(t *testing.T) {
	store := New()
	doc1 := crwDocs{{ID: 1}}
	doc2 := crwDocs{{ID: 2}, {ID: 3}}
	store.Append(doc1)
	store.Append(doc2)

	got := len(store.docs)
	want := len(doc1) + len(doc2)
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestStore_Search(t *testing.T) {
	param := "B"

	type fields struct {
		counter int
		docs    crwDocs
		ind     map[uint32][]int
	}
	type args struct {
		param *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "with filled docs",
			fields: fields{
				counter: 0,
				docs:    crwDocs{{ID: 3, URL: "golang.org", Title: "B"}, {ID: 1, URL: "golang.org", Title: "A"}, {ID: 2, URL: "golang.org", Title: "B"}},
				ind:     make(map[uint32][]int),
			},
			args: args{param: &param},
			want: []string{"2: B (golang.org)", "3: B (golang.org)"},
		},
		{
			name: "with empty docs",
			fields: fields{
				counter: 0,
				docs:    crwDocs{},
				ind:     make(map[uint32][]int),
			},
			args: args{param: &param},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				counter: tt.fields.counter,
				docs:    tt.fields.docs,
				ind:     tt.fields.ind,
			}
			s.Sort()
			s.Index()
			if got := s.Search(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmarks

func BenchmarkBinarySearch(b *testing.B) {
	data := seeds()
	store := New()
	store.Append(data)
	store.Index()
	store.Sort()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(1_000_000)
		res := store.Search(&data[n].Title)
		_ = res
	}
}

func seeds() crwDocs {
	rand.Seed(time.Now().UnixNano())
	var res crwDocs
	var cd crawler.Document

	for i := 0; i < 1_000_000; i++ {
		cd.ID = rand.Intn(1_000_000)
		cd.Title = RandStringBytesMaskImpr(10)
		res = append(res, cd)
	}

	sort.Slice(res, func(i, j int) bool { return res[i].ID < res[j].ID })
	return res
}

// Helper for strings generation. Source: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandStringBytesMaskImpr(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
