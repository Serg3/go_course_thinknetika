package storage

import (
	"bufio"
	"encoding/json"
	"go_course_thinknetika/08_search_engine/pkg/crawler"
	"io"
)

type Filestore struct{}

func New() *Filestore {
	return &Filestore{}
}

func (f *Filestore) Load(src io.Reader) ([]crawler.Document, error) {
	b, err := read(src)
	if err != nil {
		return nil, err
	}

	crwDocs := make([]crawler.Document, 0)

	err = json.Unmarshal(b, &crwDocs)
	if err != nil {
		return nil, err
	}

	return crwDocs, nil
}

func (f *Filestore) Save(src io.Writer, crwDocs []crawler.Document) error {
	b, err := json.Marshal(crwDocs)
	if err != nil {
		return err
	}

	err = write(src, b)
	if err != nil {
		return err
	}

	return nil
}

func read(r io.Reader) ([]byte, error) {
	s := bufio.NewScanner(r)
	var b []byte

	for s.Scan() {
		b = append(b, []byte(s.Text()+"\n")...)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return b, nil
}

func write(w io.Writer, b []byte) error {
	_, err := w.Write(b)
	return err
}
