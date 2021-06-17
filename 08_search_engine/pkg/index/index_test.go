package index

import (
	"go_course_thinknetika/08_search_engine/pkg/crawler"
	"reflect"
	"testing"
)

// Tests

// Table tests example
func TestStore_Empty(t *testing.T) {
	docs := crwDocs{{ID: 1}}
	docsEmpty := crwDocs{}

	type fields struct {
		counter int
		docs    crwDocs
		ind     map[uint32][]int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "with filled docs",
			fields: fields{docs: docs},
			want:   false,
		},
		{
			name:   "with empty docs",
			fields: fields{docs: docsEmpty},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				counter: tt.fields.counter,
				docs:    tt.fields.docs,
				ind:     tt.fields.ind,
			}
			if got := s.Empty(); got != tt.want {
				t.Errorf("Store.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Docs(t *testing.T) {
	store := New()
	docs := []crawler.Document{{ID: 1}}
	store.Append(docs)

	got := store.Docs()
	want := docs
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestStore_Append(t *testing.T) {
	store := New()
	doc1 := []crawler.Document{{ID: 1}}
	doc2 := []crawler.Document{{ID: 2}, {ID: 3}}
	store.Append(doc1)
	store.Append(doc2)

	got := len(store.docs)
	want := len(doc1) + len(doc2)
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestStore_Search(t *testing.T) {
	store := New()
	docs := []crawler.Document{{ID: 3, URL: "golang.org", Title: "B"}, {ID: 1, URL: "golang.org", Title: "A"}, {ID: 2, URL: "golang.org", Title: "B"}}
	want := []string{"1: B (golang.org)", "3: B (golang.org)"}
	store.Append(docs)
	store.Index()

	param := "B"
	got := store.Search(&param)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
