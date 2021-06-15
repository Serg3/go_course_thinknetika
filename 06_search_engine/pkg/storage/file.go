package storage

import (
	"bufio"
	"encoding/json"
	"go_course_thinknetika/06_search_engine/pkg/crawler"
	"io"
	"os"
)

const file = "scan_result.txt"

func Load() ([]crawler.Document, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	b, err := read(f)
	if err != nil {
		return nil, err
	}
	f.Close()

	crwDocs := make([]crawler.Document, 0)

	err = json.Unmarshal(b, &crwDocs)
	if err != nil {
		return nil, err
	}

	return crwDocs, nil
}

func Save(crwDocs []crawler.Document) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	b, err := json.Marshal(crwDocs)
	if err != nil {
		return err
	}

	err = write(f, b)
	if err != nil {
		return err
	}
	f.Close()

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
